package bdd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type sourceBinding struct {
	Kind       StepKind
	Expression string
}

func TestEveryAdapterBindingsMatchTheRecursiveCorpus(t *testing.T) {
	catalog, err := CanonicalCatalog()
	if err != nil {
		t.Fatalf("discover canonical behaviour: %v", err)
	}
	for name, bindings := range map[string][]sourceBinding{
		"owner": canonicalBindings(t),
		"e2e":   e2eBindings(t),
	} {
		if findings := validateBindings(bindings, catalog.Steps()); len(findings) != 0 {
			t.Errorf("validate %s bindings: %v", name, findings)
		}
	}
}

func TestBindingComplianceRejectsUndefinedStep(t *testing.T) {
	findings := validateBindings(nil, []Step{{Kind: WhenStep, Text: "an undefined action"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected one undefined-step finding, got %v", findings)
	}
}

func TestBindingComplianceRejectsAmbiguousStep(t *testing.T) {
	bindings := []sourceBinding{
		{Kind: ThenStep, Expression: `^the result is (.*)$`},
		{Kind: ThenStep, Expression: `^the result is good$`},
	}
	findings := validateBindings(bindings, []Step{{Kind: ThenStep, Text: "the result is good"}})
	if len(findings) == 0 || !strings.Contains(findings[0].Error(), "ambiguous") {
		t.Fatalf("expected an ambiguous-step finding, got %v", findings)
	}
}

func TestBindingComplianceRejectsUnusedBinding(t *testing.T) {
	bindings := []sourceBinding{
		{Kind: GivenStep, Expression: `^a used fixture$`},
		{Kind: GivenStep, Expression: `^an unused fixture$`},
	}
	findings := validateBindings(bindings, []Step{{Kind: GivenStep, Text: "a used fixture"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "unused") {
		t.Fatalf("expected one unused-binding finding, got %v", findings)
	}
}

func TestBindingComplianceAddedBindingFixture(t *testing.T) {
	bindings := []sourceBinding{
		{Kind: GivenStep, Expression: `^a used fixture$`},
		{Kind: GivenStep, Expression: `^an added fixture$`},
	}
	findings := validateBindings(bindings, []Step{{Kind: GivenStep, Text: "a used fixture"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "unused") {
		t.Fatalf("expected added binding to be unused, got %v", findings)
	}
}

func TestBindingComplianceEditedBindingFixture(t *testing.T) {
	bindings := []sourceBinding{{Kind: WhenStep, Expression: `^the edited action$`}}
	findings := validateBindings(bindings, []Step{{Kind: WhenStep, Text: "the original action"}})
	if len(findings) != 2 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected edited binding mismatch, got %v", findings)
	}
}

func TestBindingComplianceRenamedBindingFixture(t *testing.T) {
	bindings := []sourceBinding{{Kind: ThenStep, Expression: `^the renamed result$`}}
	findings := validateBindings(bindings, []Step{{Kind: ThenStep, Text: "the old result"}})
	if len(findings) != 2 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected renamed binding mismatch, got %v", findings)
	}
}

func TestBindingComplianceDeletedBindingFixture(t *testing.T) {
	findings := validateBindings(nil, []Step{{Kind: GivenStep, Text: "a deleted fixture"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected deleted binding mismatch, got %v", findings)
	}
}

func canonicalBindings(t *testing.T) []sourceBinding {
	t.Helper()
	root, err := moduleRoot()
	if err != nil {
		t.Fatalf("locate module root: %v", err)
	}
	return bindingsAt(t, filepath.Join(root, "tests", "bdd", "bindings.go"))
}

func e2eBindings(t *testing.T) []sourceBinding {
	t.Helper()
	root, err := moduleRoot()
	if err != nil {
		t.Fatalf("locate module root: %v", err)
	}
	return bindingsAt(t, filepath.Join(root, "..", "badakmini-cli-e2e", "bindings_test.go"))
}

func bindingsAt(t *testing.T, path string) []sourceBinding {
	t.Helper()
	bindings, err := discoverGodogBindings(path)
	if err != nil {
		t.Fatalf("discover direct Godog bindings at %s: %v", path, err)
	}
	return bindings
}

func discoverGodogBindings(path string) ([]sourceBinding, error) {
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parse Godog bindings: %w", err)
	}
	initializer, err := findScenarioInitializer(file)
	if err != nil {
		return nil, err
	}
	receiver := initializer.Type.Params.List[0].Names[0].Name
	bindings := make([]sourceBinding, 0, len(initializer.Body.List))
	for _, statement := range initializer.Body.List {
		binding, bindingErr := directGodogBinding(statement, receiver)
		if bindingErr != nil {
			return nil, bindingErr
		}
		bindings = append(bindings, binding)
	}
	if len(bindings) == 0 {
		return nil, errors.New("initialize scenario has no direct Godog step definitions")
	}
	return bindings, nil
}

func directGodogBinding(statement ast.Stmt, receiver string) (sourceBinding, error) {
	call, err := directGodogCall(statement)
	if err != nil {
		return sourceBinding{}, err
	}
	selector, isSelector := call.Fun.(*ast.SelectorExpr)
	if !isSelector {
		return sourceBinding{}, errors.New("initialize scenario must call Godog directly")
	}
	identifier, isIdentifier := selector.X.(*ast.Ident)
	kind, recognized := registrationKind(selector.Sel.Name)
	if !isIdentifier || identifier.Name != receiver || !recognized {
		return sourceBinding{}, errors.New("initialize scenario must call Godog Given, When, or Then directly")
	}
	if len(call.Args) < 2 {
		return sourceBinding{}, fmt.Errorf("Godog %s registration requires an expression and handler", selector.Sel.Name)
	}
	literal, isLiteral := call.Args[0].(*ast.BasicLit)
	if !isLiteral || literal.Kind != token.STRING {
		return sourceBinding{}, fmt.Errorf("Godog %s expression must be a direct string literal", selector.Sel.Name)
	}
	expression, err := strconv.Unquote(literal.Value)
	if err != nil {
		return sourceBinding{}, fmt.Errorf("decode Godog %s expression: %w", selector.Sel.Name, err)
	}
	return sourceBinding{Kind: kind, Expression: expression}, nil
}

func directGodogCall(statement ast.Stmt) (*ast.CallExpr, error) {
	expressionStatement, isExpression := statement.(*ast.ExprStmt)
	if !isExpression {
		return nil, errors.New("initialize scenario must contain only direct Godog registrations")
	}
	call, isCall := expressionStatement.X.(*ast.CallExpr)
	if !isCall {
		return nil, errors.New("initialize scenario must contain only direct Godog registrations")
	}
	return call, nil
}

func findScenarioInitializer(file *ast.File) (*ast.FuncDecl, error) {
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Name.Name != "InitializeScenario" {
			continue
		}
		if function.Type.Params == nil || len(function.Type.Params.List) != 1 ||
			len(function.Type.Params.List[0].Names) != 1 {
			return nil, errors.New("initialize scenario must receive one named Godog scenario context")
		}
		return function, nil
	}
	return nil, errors.New("bindings.go must define InitializeScenario")
}

func registrationKind(method string) (StepKind, bool) {
	switch method {
	case "Given":
		return GivenStep, true
	case "When":
		return WhenStep, true
	case "Then":
		return ThenStep, true
	default:
		return "", false
	}
}

func validateBindings(bindings []sourceBinding, steps []Step) []error {
	used := make([]bool, len(bindings))
	var findings []error
	for _, step := range steps {
		matches, matchFindings := matchingBindings(bindings, step)
		findings = append(findings, matchFindings...)
		switch len(matches) {
		case 0:
			findings = append(findings, fmt.Errorf("undefined %s step %q", step.Kind, step.Text))
		case 1:
			used[matches[0]] = true
		default:
			findings = append(findings, fmt.Errorf("ambiguous %s step %q", step.Kind, step.Text))
		}
	}
	for index, binding := range bindings {
		if !used[index] {
			findings = append(findings, fmt.Errorf("unused %s binding %q", binding.Kind, binding.Expression))
		}
	}
	return findings
}

func matchingBindings(bindings []sourceBinding, step Step) ([]int, []error) {
	matches := make([]int, 0, 1)
	var findings []error
	for index, binding := range bindings {
		if binding.Kind != step.Kind {
			continue
		}
		matched, err := regexp.MatchString(binding.Expression, step.Text)
		if err != nil {
			findings = append(findings, fmt.Errorf("invalid binding %q: %w", binding.Expression, err))
			continue
		}
		if matched {
			matches = append(matches, index)
		}
	}
	return matches, findings
}
