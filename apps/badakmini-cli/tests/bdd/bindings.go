package bdd

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/cucumber/godog"
)

// Definition is one typed binding and its scenario-scoped implementation.
type Definition struct {
	Kind       StepKind
	Expression string
	Handler    func(*State) any
}

// Registry is the single binding catalog shared by every adapter.
type Registry struct {
	definitions []Definition
}

// NewRegistry creates a binding catalog from explicit definitions.
func NewRegistry(definitions ...Definition) Registry {
	return Registry{definitions: append([]Definition(nil), definitions...)}
}

// CanonicalRegistry owns the step vocabulary shared by every Badak Mini adapter.
//
//nolint:funlen // One explicit catalog makes duplicate and missing Gherkin expressions reviewable.
func CanonicalRegistry() Registry {
	return NewRegistry(
		Definition{
			Kind:       GivenStep,
			Expression: "repository discovery would fail",
			Handler:    prepareFixture("repository-discovery-fails"),
		},
		Definition{
			Kind:       WhenStep,
			Expression: `Badak Mini runs with "--help"`,
			Handler:    invokeCommand("--help"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds and prints usage",
			Handler:    expectResult(0, "Usage:", ""),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository whose governance documents fit the word limit",
			Handler:    prepareFixture("governance-documents-fit"),
		},
		Definition{
			Kind:       WhenStep,
			Expression: "Badak Mini runs instruction-size validation",
			Handler:    invokeCommand("harness", "instruction-size", "validate"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds with the word-limit confirmation",
			Handler:    expectResult(0, "Governance word counts are within", ""),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository with an oversized agent instruction file",
			Handler:    prepareFixture("oversized-agent-instruction"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command fails with the oversized document diagnostic",
			Handler:    expectResult(1, "", "AGENTS.md contains 501 words"),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository whose tracked Markdown links resolve",
			Handler:    prepareFixture("tracked-markdown-links-resolve"),
		},
		Definition{
			Kind:       WhenStep,
			Expression: "Badak Mini runs Markdown-link validation",
			Handler:    invokeCommand("harness", "markdown-links", "validate"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds with the link confirmation",
			Handler:    expectResult(0, "Repository-local Markdown links are valid", ""),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository with a broken tracked Markdown link",
			Handler:    prepareFixture("broken-tracked-markdown-link"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command fails with the missing-target diagnostic",
			Handler:    expectResult(1, "", "targets a file that does not exist"),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository whose harness capabilities match",
			Handler:    prepareFixture("harness-capabilities-match"),
		},
		Definition{
			Kind:       WhenStep,
			Expression: "Badak Mini runs capability-parity validation",
			Handler:    invokeCommand("harness", "capability-parity", "validate"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds with the parity confirmation",
			Handler:    expectResult(0, "Every harness exposes the same", ""),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository with a harness missing a shared subagent",
			Handler:    prepareFixture("harness-missing-shared-subagent"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command fails with the parity diagnostic",
			Handler:    expectResult(1, "", "subagent parity"),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository with a staged rule-bearing file",
			Handler:    prepareFixture("staged-rule-bearing-file"),
		},
		Definition{
			Kind:       WhenStep,
			Expression: "Badak Mini runs staged rule-change detection",
			Handler:    invokeCommand("harness", "rule-change", "validate"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds with the rules-propagation notice",
			Handler:    expectResult(0, "repo-governance/workflows/rules-propagation.md", ""),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a repository with only an ordinary staged file",
			Handler:    prepareFixture("ordinary-staged-file"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds without output",
			Handler:    expectResult(0, "", ""),
		},
		Definition{
			Kind:       GivenStep,
			Expression: "a pre-edit payload for a harness instruction file",
			Handler:    prepareFixture("harness-instruction-pre-edit"),
		},
		Definition{
			Kind:       WhenStep,
			Expression: "Badak Mini runs hook rule-change detection",
			Handler:    invokeCommand("harness", "rule-change", "hook"),
		},
		Definition{
			Kind:       ThenStep,
			Expression: "the command succeeds with both workflow notices",
			Handler: expectStdoutContains(
				"repo-governance/workflows/rules-propagation.md",
				"repo-governance/workflows/harness-alignment.md",
			),
		},
	)
}

func prepareFixture(fixture string) func(*State) any {
	return func(state *State) any {
		return func() error { return state.Driver().Prepare(context.Background(), fixture) }
	}
}

func invokeCommand(arguments ...string) func(*State) any {
	return func(state *State) any {
		return func() error { return state.Driver().Invoke(context.Background(), arguments) }
	}
}

func expectResult(exitCode int, stdoutContains, stderrContains string) func(*State) any {
	return func(state *State) any {
		return func() error {
			result := state.Driver().Result()
			if result.ExitCode != exitCode || !matchesOutput(result.Stdout, stdoutContains) ||
				!matchesOutput(result.Stderr, stderrContains) {
				return fmt.Errorf(
					"expected exit %d, stdout %q, stderr %q; got %#v",
					exitCode,
					stdoutContains,
					stderrContains,
					result,
				)
			}
			return nil
		}
	}
}

func matchesOutput(output, expected string) bool {
	if expected == "" {
		return output == ""
	}
	return strings.Contains(output, expected)
}

func expectStdoutContains(expected ...string) func(*State) any {
	return func(state *State) any {
		return func() error {
			result := state.Driver().Result()
			if result.ExitCode != 0 || result.Stderr != "" {
				return fmt.Errorf("expected successful stdout-only result, got %#v", result)
			}
			for _, value := range expected {
				if !strings.Contains(result.Stdout, value) {
					return fmt.Errorf("expected stdout to contain %q, got %#v", value, result)
				}
			}
			return nil
		}
	}
}

// Definitions returns a defensive copy for static compliance checks.
func (registry Registry) Definitions() []Definition {
	return append([]Definition(nil), registry.definitions...)
}

// Register installs every binding against scenario-scoped state.
func (registry Registry) Register(context *godog.ScenarioContext, state **State) {
	for _, definition := range registry.definitions {
		current := definition
		context.Step(regexp.MustCompile(current.Expression), func() error {
			handler, ok := current.Handler(*state).(func() error)
			if !ok {
				return fmt.Errorf("binding %q has an unsupported handler", current.Expression)
			}

			return handler()
		})
	}
}

// Validate proves every canonical step resolves once and every binding is used.
func (registry Registry) Validate(steps []Step) []error {
	used := make([]bool, len(registry.definitions))
	var findings []error
	for _, step := range steps {
		matches, matchFindings := registry.matches(step)
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
	for index, definition := range registry.definitions {
		if !used[index] {
			findings = append(findings, fmt.Errorf("unused %s binding %q", definition.Kind, definition.Expression))
		}
	}

	return findings
}

func (registry Registry) matches(step Step) ([]int, []error) {
	matches := make([]int, 0, 1)
	var findings []error
	for index, definition := range registry.definitions {
		if definition.Kind != step.Kind {
			continue
		}
		matched, err := regexp.MatchString("^(?:"+definition.Expression+")$", step.Text)
		if err != nil {
			findings = append(findings, fmt.Errorf("invalid binding %q: %w", definition.Expression, err))
			continue
		}
		if matched {
			matches = append(matches, index)
		}
	}

	return matches, findings
}
