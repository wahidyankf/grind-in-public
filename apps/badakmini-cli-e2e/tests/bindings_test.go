package e2e_test

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
)

type scenarioStateKey struct{}

type scenarioState struct {
	driver *processDriver
}

// InitializeScenario registers the process E2E step definitions directly with Godog.
//
//nolint:funlen // One explicit catalog makes duplicate and missing Gherkin expressions reviewable.
func InitializeScenario(context *godog.ScenarioContext) {
	context.Given(`^repository discovery would fail$`, prepareFixture("repository-discovery-fails"))
	context.When(`^Badak Mini runs with "--help"$`, invokeCommand("--help"))
	context.Then(`^the command succeeds and prints usage$`, expectResult(0, "Usage:", ""))
	context.Given(
		`^a repository whose governance documents fit the word limit$`,
		prepareFixture("governance-documents-fit"),
	)
	context.When(
		`^Badak Mini runs instruction-size validation$`,
		invokeCommand("harness", "instruction-size", "validate"),
	)
	context.Then(
		`^the command succeeds with the word-limit confirmation$`,
		expectResult(0, "Governance word counts are within", ""),
	)
	context.Given(
		`^a repository with an oversized agent instruction file$`,
		prepareFixture("oversized-agent-instruction"),
	)
	context.Then(
		`^the command fails with the oversized document diagnostic$`,
		expectResult(1, "", "AGENTS.md contains 501 words"),
	)
	context.Given(
		`^a repository whose tracked Markdown links resolve$`,
		prepareFixture("tracked-markdown-links-resolve"),
	)
	context.When(
		`^Badak Mini runs Markdown-link validation$`,
		invokeCommand("harness", "markdown-links", "validate"),
	)
	context.Then(
		`^the command succeeds with the link confirmation$`,
		expectResult(0, "Repository-local Markdown links are valid", ""),
	)
	context.Given(
		`^a repository with a broken tracked Markdown link$`,
		prepareFixture("broken-tracked-markdown-link"),
	)
	context.Then(
		`^the command fails with the missing-target diagnostic$`,
		expectResult(1, "", "targets a file that does not exist"),
	)
	context.Given(
		`^a repository whose harness capabilities match$`,
		prepareFixture("harness-capabilities-match"),
	)
	context.When(
		`^Badak Mini runs capability-parity validation$`,
		invokeCommand("harness", "capability-parity", "validate"),
	)
	context.Then(
		`^the command succeeds with the parity confirmation$`,
		expectResult(0, "Every harness exposes the same", ""),
	)
	context.Given(
		`^a repository with a harness missing a shared subagent$`,
		prepareFixture("harness-missing-shared-subagent"),
	)
	context.Then(
		`^the command fails with the parity diagnostic$`,
		expectResult(1, "", "subagent parity"),
	)
	context.Given(
		`^a repository with a staged rule-bearing file$`,
		prepareFixture("staged-rule-bearing-file"),
	)
	context.When(
		`^Badak Mini runs staged rule-change detection$`,
		invokeCommand("harness", "rule-change", "validate"),
	)
	context.Then(
		`^the command succeeds with the rules-propagation notice$`,
		expectResult(0, "repo-governance/workflows/rules-propagation.md", ""),
	)
	context.Given(
		`^a repository with only an ordinary staged file$`,
		prepareFixture("ordinary-staged-file"),
	)
	context.Then(`^the command succeeds without output$`, expectResult(0, "", ""))
	context.Given(
		`^a pre-edit payload for a harness instruction file$`,
		prepareFixture("harness-instruction-pre-edit"),
	)
	context.When(
		`^Badak Mini runs hook rule-change detection$`,
		invokeCommand("harness", "rule-change", "hook"),
	)
	context.Then(
		`^the command succeeds with both workflow notices$`,
		expectStdoutContains(
			"repo-governance/workflows/rules-propagation.md",
			"repo-governance/workflows/harness-alignment.md",
		),
	)
}

func contextWithState(ctx context.Context, state *scenarioState) context.Context {
	return context.WithValue(ctx, scenarioStateKey{}, state)
}

func stateFromContext(ctx context.Context) (*scenarioState, error) {
	state, ok := ctx.Value(scenarioStateKey{}).(*scenarioState)
	if !ok {
		return nil, errors.New("godog scenario context has no Badak Mini E2E state")
	}
	return state, nil
}

func prepareFixture(fixture string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
		return state.driver.Prepare(ctx, fixture)
	}
}

func invokeCommand(arguments ...string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
		return state.driver.Invoke(ctx, arguments)
	}
}

func expectResult(exitCode int, stdoutContains, stderrContains string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
		result := state.driver.Result()
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

func matchesOutput(output, expected string) bool {
	if expected == "" {
		return output == ""
	}
	return strings.Contains(output, expected)
}

func expectStdoutContains(expected ...string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
		result := state.driver.Result()
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
