package bdd

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
)

// InitializeScenario registers Badak Mini's step definitions directly with Godog.
//
//nolint:funlen,varnamelen // The explicit catalog exposes drift; sc consistently means scenario context.
func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^repository discovery would fail$`, prepareFixture("repository-discovery-fails"))
	sc.When(`^Badak Mini runs with "([^"]+)"$`, invokeCommandLine)
	sc.Then(`^the command succeeds and prints usage$`, expectResult(0, "Usage:", ""))
	sc.Given(
		`^a repository whose governance documents fit the word limit$`,
		prepareFixture("governance-documents-fit"),
	)
	sc.When(
		`^Badak Mini runs instruction-size validation$`,
		invokeCommand("harness", "instruction-size", "validate"),
	)
	sc.Then(
		`^the command succeeds with the word-limit confirmation$`,
		expectResult(0, "Governance word counts are within", ""),
	)
	sc.Given(
		`^a repository with an oversized agent instruction file$`,
		prepareFixture("oversized-agent-instruction"),
	)
	sc.Then(
		`^the command fails with the oversized document diagnostic$`,
		expectResult(1, "", "AGENTS.md contains"),
	)
	sc.Given(
		`^a repository whose tracked Markdown links resolve$`,
		prepareFixture("tracked-markdown-links-resolve"),
	)
	sc.When(
		`^Badak Mini runs Markdown-link validation$`,
		invokeCommand("harness", "markdown-links", "validate"),
	)
	sc.Then(
		`^the command succeeds with the link confirmation$`,
		expectResult(0, "Repository-local Markdown links are valid", ""),
	)
	sc.Given(
		`^a repository with a broken tracked Markdown link$`,
		prepareFixture("broken-tracked-markdown-link"),
	)
	sc.Then(
		`^the command fails with the missing-target diagnostic$`,
		expectResult(1, "", "targets a file that does not exist"),
	)
	sc.Given(
		`^a repository whose canonical harness contract matches$`,
		prepareFixture("canonical-harness-contract-matches"),
	)
	sc.When(
		`^Badak Mini runs capability-parity validation$`,
		invokeCommand("harness", "capability-parity", "validate"),
	)
	sc.Then(
		`^the command succeeds with canonical counts and a digest$`,
		expectStdoutContains("3 harnesses", "1 skill", "1 agent", "sha256:"),
	)
	sc.Given(
		`^a canonical harness contract with a missing Codex agent adapter$`,
		prepareFixture("missing-codex-agent-adapter"),
	)
	sc.Then(
		`^the command fails with a missing-agent-adapter diagnostic$`,
		expectResult(1, "", "missing-agent-adapter"),
	)
	sc.Given(
		`^a canonical harness contract with an instruction overlay$`,
		prepareFixture("instruction-overlay"),
	)
	sc.Then(
		`^the command fails with an unexpected-instruction-source diagnostic$`,
		expectResult(1, "", "unexpected-instruction-source"),
	)
	sc.Given(
		`^a canonical harness contract with a stale Claude skill adapter$`,
		prepareFixture("stale-claude-skill-adapter"),
	)
	sc.Then(
		`^the command fails with a skill-content-divergence diagnostic$`,
		expectResult(1, "", "skill-content-divergence"),
	)
	sc.Given(
		`^a canonical harness contract with weakened opencode permissions$`,
		prepareFixture("weakened-opencode-permissions"),
	)
	sc.Then(
		`^the command fails with an agent-semantic-divergence diagnostic$`,
		expectResult(1, "", "agent-semantic-divergence"),
	)
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
