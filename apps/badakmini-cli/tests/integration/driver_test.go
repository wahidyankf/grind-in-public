package integration_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

type behaviourDriver struct {
	testing *testing.T
	root    string
	stdout  bytes.Buffer
	stderr  bytes.Buffer
	runtime cli.Runtime
	result  bdd.Result
}

func newBehaviourDriver(t *testing.T) *behaviourDriver {
	t.Helper()
	driver := &behaviourDriver{testing: t, root: t.TempDir()}
	driver.resetRuntime()
	return driver
}

//nolint:funcorder // Constructor setup stays beside newBehaviourDriver for the fixture lifecycle.
func (driver *behaviourDriver) resetRuntime() {
	driver.stdout.Reset()
	driver.stderr.Reset()
	driver.result = bdd.Result{}
	driver.runtime = cli.Runtime{
		Stdin:              strings.NewReader(""),
		Stdout:             &driver.stdout,
		Stderr:             &driver.stderr,
		FindRepositoryRoot: func() (string, error) { return driver.root, nil },
		CheckGovernance:    func(root string) ([]governance.Finding, error) { return governance.CheckFS(os.DirFS(root)) },
		CheckMarkdownLinks: func(root string) ([]markdownlinks.Finding, error) {
			return markdownlinks.Check(root, realMarkdownRuntime())
		},
		ListStagedPaths: realStagedPaths,
		CheckParity:     func(root string) (parity.Report, error) { return parity.CheckFS(os.DirFS(root)) },
	}
}

//nolint:cyclop,funlen // One switch keeps canonical fixture selection visible and exhaustive.
func (driver *behaviourDriver) Prepare(_ context.Context, fixture string) error {
	driver.resetRuntime()
	switch fixture {
	case "repository-discovery-fails":
		driver.runtime.FindRepositoryRoot = func() (string, error) {
			return "", errors.New("repository discovery failed")
		}
	case "governance-documents-fit":
		driver.writeGovernance("short")
	case "oversized-agent-instruction":
		driver.writeGovernance(strings.Repeat("word ", governance.MaxWords+1))
	case "tracked-markdown-links-resolve":
		runMarkdownGit(driver.testing, driver.root, "init", "--quiet")
		writeMarkdownFile(driver.testing, driver.root, "README.md", "[Guide](docs/guide.md)\n")
		writeMarkdownFile(driver.testing, driver.root, "docs/guide.md", "# Guide\n")
		runMarkdownGit(driver.testing, driver.root, "add", "README.md", "docs/guide.md")
	case "broken-tracked-markdown-link":
		runMarkdownGit(driver.testing, driver.root, "init", "--quiet")
		writeMarkdownFile(driver.testing, driver.root, "README.md", "[Missing](missing.md)\n")
		runMarkdownGit(driver.testing, driver.root, "add", "README.md")
	case "canonical-harness-contract-matches":
		driver.writeCanonicalHarnessContract()
	case "missing-codex-agent-adapter":
		driver.writeCanonicalHarnessContract()
		if err := os.Remove(driver.root + "/.codex/agents/review.toml"); err != nil {
			return fmt.Errorf("remove Codex adapter fixture: %w", err)
		}
	case "instruction-overlay":
		driver.writeCanonicalHarnessContract()
		writeParityFile(driver.testing, driver.root, "nested/AGENTS.md", "overlay")
	case "stale-claude-skill-adapter":
		driver.writeCanonicalHarnessContract()
		writeParityFile(driver.testing, driver.root, ".claude/skills/review/SKILL.md", "stale")
	case "weakened-opencode-permissions":
		driver.writeCanonicalHarnessContract()
		weakened := strings.Replace(opencodeAgentFixture, "edit: deny", "edit: allow", 1)
		writeParityFile(driver.testing, driver.root, ".opencode/agents/review.md", weakened)
	case "staged-rule-bearing-file":
		driver.stageFile("repo-governance/development/testing-policy.md")
	case "ordinary-staged-file":
		driver.stageFile("README.md")
	case "harness-instruction-pre-edit":
		driver.runtime.Stdin = strings.NewReader(`{"tool_input":{"file_path":"AGENTS.md"}}`)
	default:
		return errors.New("unsupported integration fixture: " + fixture)
	}
	return nil
}

func (driver *behaviourDriver) Invoke(ctx context.Context, arguments []string) error {
	driver.result.ExitCode = cli.Run(ctx, driver.runtime, arguments)
	driver.result.Stdout = driver.stdout.String()
	driver.result.Stderr = driver.stderr.String()
	return nil
}

func (driver *behaviourDriver) Result() bdd.Result {
	return driver.result
}

func (driver *behaviourDriver) writeGovernance(agents string) {
	writeGovernanceFile(driver.testing, driver.root, "AGENTS.md", agents)
	writeGovernanceFile(driver.testing, driver.root, "CLAUDE.md", "short")
	writeGovernanceFile(driver.testing, driver.root, "RTK.md", "short")
	writeGovernanceFile(driver.testing, driver.root, "repo-governance/policy.md", "short")
}

func (driver *behaviourDriver) writeCanonicalHarnessContract() {
	writeCanonicalHarnessFixture(driver.testing, driver.root)
}

func (driver *behaviourDriver) stageFile(path string) {
	runRuleChangeGit(driver.testing, driver.root, "init", "--quiet")
	writeRuleChangeFile(driver.testing, driver.root, path, "fixture")
	runRuleChangeGit(driver.testing, driver.root, "add", "--", path)
}
