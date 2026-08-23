package integration_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

type behaviorDriver struct {
	testing *testing.T
	root    string
	stdout  bytes.Buffer
	stderr  bytes.Buffer
	runtime cli.Runtime
	result  bdd.Result
}

func newBehaviorDriver(t *testing.T) *behaviorDriver {
	t.Helper()
	driver := &behaviorDriver{testing: t, root: t.TempDir()}
	driver.resetRuntime()
	return driver
}

//nolint:funcorder // Constructor setup stays beside newBehaviorDriver for the fixture lifecycle.
func (driver *behaviorDriver) resetRuntime() {
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
		CheckParity:     func(root string) ([]parity.Finding, error) { return parity.CheckFS(os.DirFS(root)) },
	}
}

//nolint:cyclop // One switch keeps canonical fixture selection visible and exhaustive.
func (driver *behaviorDriver) Prepare(_ context.Context, fixture string) error {
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
	case "harness-capabilities-match":
		driver.writeMatchingHarnesses()
	case "harness-missing-shared-subagent":
		writeParityFile(driver.testing, driver.root, ".claude/agents/review.md", "fixture")
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

func (driver *behaviorDriver) Invoke(ctx context.Context, arguments []string) error {
	driver.result.ExitCode = cli.Run(ctx, driver.runtime, arguments)
	driver.result.Stdout = driver.stdout.String()
	driver.result.Stderr = driver.stderr.String()
	return nil
}

func (driver *behaviorDriver) Result() bdd.Result {
	return driver.result
}

func (driver *behaviorDriver) writeGovernance(agents string) {
	writeGovernanceFile(driver.testing, driver.root, "AGENTS.md", agents)
	writeGovernanceFile(driver.testing, driver.root, "CLAUDE.md", "short")
	writeGovernanceFile(driver.testing, driver.root, "repo-governance/policy.md", "short")
}

func (driver *behaviorDriver) writeMatchingHarnesses() {
	for _, path := range []string{
		".claude/agents/review.md",
		".codex/agents/review.toml",
		".opencode/agents/review.md",
	} {
		writeParityFile(driver.testing, driver.root, path, "fixture")
	}
}

func (driver *behaviorDriver) stageFile(path string) {
	runRuleChangeGit(driver.testing, driver.root, "init", "--quiet")
	writeRuleChangeFile(driver.testing, driver.root, path, "fixture")
	runRuleChangeGit(driver.testing, driver.root, "add", "--", path)
}
