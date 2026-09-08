package bdd

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
)

// invalidInvocation is the exit code the CLI contract reserves for a name or
// argument shape this CLI does not have. The adapters assert the public
// contract, so the number is stated here rather than read from the
// implementation it is meant to hold to account.
const invalidInvocation = 2

// InitializeScenario registers Badak Mini's step definitions directly with Godog.
//
//nolint:varnamelen // sc consistently means scenario context.
func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^repository discovery would fail$`, prepareFixture("repository-discovery-fails"))
	sc.When(`^Badak Mini runs with "([^"]+)"$`, invokeCommandLine)
	sc.Then(`^the command succeeds and prints usage$`, expectResult(0, "Usage:", ""))
	sc.Then(`^the command reports an invalid invocation$`, expectResult(invalidInvocation, "", "Usage:"))
	sc.Given(
		`^a repository with a staged rule-bearing file$`,
		prepareFixture("staged-rule-bearing-file"),
	)
	sc.When(
		`^Badak Mini runs staged rule-change detection$`,
		invokeCommand("harness", "rule-change", "validate"),
	)
	sc.Then(
		`^the command succeeds with the automatically triggered rules-propagation workflow$`,
		expectResult(0, "repo-governance/workflows/rules/rules-propagation.md", ""),
	)
	sc.Given(
		`^a repository with only an ordinary staged file$`,
		prepareFixture("ordinary-staged-file"),
	)
	sc.Then(`^the command succeeds without output$`, expectResult(0, "", ""))
	sc.Given(
		`^a pre-edit payload for a harness instruction file$`,
		prepareFixture("harness-instruction-pre-edit"),
	)
	sc.When(
		`^Badak Mini runs hook rule-change detection$`,
		invokeCommand("harness", "rule-change", "hook"),
	)
	sc.Then(
		`^the command succeeds with both workflow notices$`,
		expectStdoutContains(
			"repo-governance/workflows/rules/rules-propagation.md",
			"repo-governance/workflows/harness-alignment.md",
		),
	)
}

func invokeCommandLine(ctx context.Context, arguments string) error {
	return invokeCommand(strings.Fields(arguments)...)(ctx)
}

func prepareFixture(fixture string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
		return state.Driver().Prepare(ctx, fixture)
	}
}

func invokeCommand(arguments ...string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
		return state.Driver().Invoke(ctx, arguments)
	}
}

func expectResult(exitCode int, stdoutContains, stderrContains string) func(context.Context) error {
	return func(ctx context.Context) error {
		state, err := stateFromContext(ctx)
		if err != nil {
			return err
		}
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
