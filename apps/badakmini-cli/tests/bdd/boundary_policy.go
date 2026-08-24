package bdd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

var forbiddenUnitImports = []string{
	"net",
	"net/http",
	"net/http/httptest",
	"os",
	"os/exec",
}

var forbiddenUnitMethods = []string{"Chdir", "Setenv", "TempDir"}

var allowedUnitImportsByFile = map[string][]string{
	// Adapter parity statically reads Nx target inputs to prove E2E changes invalidate the behavior gate.
	"adapter_parity_test.go": {"os"},
}

var integrationOwners = map[string]string{
	"cli_test.go":           "TestIntegrationCLI",
	"governance_test.go":    "TestIntegrationGovernance",
	"markdownlinks_test.go": "TestIntegrationMarkdownLinks",
	"parity_test.go":        "TestIntegrationParity",
	"rulechange_test.go":    "TestIntegrationRuleChange",
}

var forbiddenIntegrationImports = []string{
	"net",
	"net/http",
	"net/http/httptest",
	"net/rpc",
}

var loopbackMarkers = []string{"127.0.0.1", "[::1]", "localhost"}

func unitBoundaryFindings() ([]string, error) {
	root, err := moduleRoot()
	if err != nil {
		return nil, err
	}

	var findings []string
	for _, directory := range []string{"cmd", "internal", "tests/bdd", "tests/unit"} {
		directoryFindings, err := inspectTestFiles(filepath.Join(root, directory), inspectUnitTest)
		if err != nil {
			return nil, err
		}
		findings = append(findings, directoryFindings...)
	}
	return findings, nil
}

func integrationBoundaryFindings() ([]string, error) {
	root, err := moduleRoot()
	if err != nil {
		return nil, err
	}
	integrationRoot := filepath.Join(root, "tests", "integration")
	seenOwners := make(map[string]bool, len(integrationOwners))
	findings, err := inspectTestFiles(integrationRoot, func(path string, file *ast.File) []string {
		return inspectIntegrationTest(path, file, seenOwners)
	})
	if err != nil {
		return nil, err
	}
	for filename, prefix := range integrationOwners {
		if !seenOwners[filename] {
			findings = append(findings, fmt.Sprintf("tests/integration/%s must own at least one %s test", filename, prefix))
		}
	}
	return findings, nil
}

//nolint:cyclop // The policy intentionally lists every process-boundary prohibition in one audit.
func e2eBoundaryFindings() ([]string, error) {
	root, err := moduleRoot()
	if err != nil {
		return nil, err
	}
	e2eRoot := filepath.Join(root, "tests", "e2e")
	var findings []string
	err = filepath.WalkDir(e2eRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		if filepath.Ext(entry.Name()) == ".feature" {
			findings = append(findings, path+" must consume canonical specs instead of owning a feature")
			return nil
		}
		if !strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		parsed, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if parseErr != nil {
			return fmt.Errorf("parse %s: %w", path, parseErr)
		}
		for _, imported := range parsed.Imports {
			importPath, unquoteErr := strconv.Unquote(imported.Path.Value)
			if unquoteErr != nil {
				continue
			}
			if strings.Contains(importPath, "/apps/badakmini-cli/internal/") {
				findings = append(findings, fmt.Sprintf("%s imports owner internal package %q", path, importPath))
			}
			if importPath == "os/exec" && filepath.Base(path) != "driver_test.go" {
				findings = append(findings, path+" imports process boundary outside driver_test.go")
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("inspect E2E tests below %s: %w", e2eRoot, err)
	}
	return findings, nil
}

func inspectTestFiles(root string, inspect func(string, *ast.File) []string) ([]string, error) {
	var findings []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		parsed, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		findings = append(findings, inspect(path, parsed)...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("inspect tests below %s: %w", root, err)
	}
	return findings, nil
}

func inspectUnitTest(path string, file *ast.File) []string {
	var findings []string
	for _, imported := range file.Imports {
		importPath, err := strconv.Unquote(imported.Path.Value)
		if err == nil && slices.Contains(forbiddenUnitImports, importPath) &&
			!slices.Contains(allowedUnitImportsByFile[filepath.Base(path)], importPath) {
			findings = append(findings, fmt.Sprintf("%s imports system boundary %q", path, importPath))
		}
	}
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if ok && slices.Contains(forbiddenUnitMethods, selector.Sel.Name) {
			findings = append(findings, fmt.Sprintf("%s calls testing system boundary %s", path, selector.Sel.Name))
		}
		return true
	})
	return findings
}

func inspectIntegrationTest(path string, file *ast.File, seenOwners map[string]bool) []string {
	findings := inspectForbiddenImports(path, file, forbiddenIntegrationImports, "network")
	findings = append(findings, inspectOwnerTests(path, file, seenOwners)...)
	findings = append(findings, inspectLoopbackMarkers(path, file)...)
	return findings
}

func inspectOwnerTests(path string, file *ast.File, seenOwners map[string]bool) []string {
	var findings []string
	filename := filepath.Base(path)
	prefix, requiredOwner := integrationOwners[filename]
	if !requiredOwner {
		return nil
	}
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || !strings.HasPrefix(function.Name.Name, "Test") {
			continue
		}
		if !strings.HasPrefix(function.Name.Name, prefix) {
			findings = append(findings, fmt.Sprintf("%s test %s must use owner prefix %s", path, function.Name.Name, prefix))
			continue
		}
		seenOwners[filename] = true
	}
	return findings
}

func inspectLoopbackMarkers(path string, file *ast.File) []string {
	var findings []string
	ast.Inspect(file, func(node ast.Node) bool {
		literal, ok := node.(*ast.BasicLit)
		if !ok || literal.Kind != token.STRING {
			return true
		}
		value, err := strconv.Unquote(literal.Value)
		if err != nil {
			return true
		}
		for _, marker := range loopbackMarkers {
			if strings.Contains(strings.ToLower(value), marker) {
				findings = append(findings, fmt.Sprintf("%s contains loopback network marker %q", path, marker))
			}
		}
		return true
	})
	return findings
}

func inspectForbiddenImports(path string, file *ast.File, forbidden []string, boundary string) []string {
	var findings []string
	for _, imported := range file.Imports {
		importPath, err := strconv.Unquote(imported.Path.Value)
		if err == nil && slices.Contains(forbidden, importPath) {
			findings = append(findings, fmt.Sprintf("%s imports %s boundary %q", path, boundary, importPath))
		}
	}
	return findings
}

func moduleRoot() (string, error) {
	_, sourcePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("locate boundary policy source")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(sourcePath), "..", "..")), nil
}
