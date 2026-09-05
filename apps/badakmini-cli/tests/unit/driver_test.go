package unit_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

// driver executes command orchestration with injected collaborators. Harness
// parity scenarios use an in-memory filesystem so they exercise the production
// validator without crossing the unit boundary.
type driver struct {
	runtime    cli.Runtime
	fileSystem fstest.MapFS
	stdout     bytes.Buffer
	stderr     bytes.Buffer
	result     bdd.Result
}

func newDriver() *driver {
	driver := &driver{fileSystem: unitCanonicalHarnessContract()}
	driver.runtime = cli.Runtime{
		Stdin:  bytes.NewReader(nil),
		Stdout: &driver.stdout,
		Stderr: &driver.stderr,
		FindRepositoryRoot: func() (string, error) {
			return "repository", nil
		},
		CheckGovernance:    func(string) ([]governance.Finding, error) { return nil, nil },
		CheckMarkdownLinks: func(string) ([]markdownlinks.Finding, error) { return nil, nil },
		ListStagedPaths:    func(string) ([]string, error) { return nil, nil },
		CheckParity: func(string) (parity.Report, error) {
			return parity.CheckFS(driver.fileSystem)
		},
	}
	return driver
}

//nolint:cyclop,funlen // One switch keeps canonical fixture selection visible and exhaustive.
func (driver *driver) Prepare(_ context.Context, fixture string) error {
	driver.stdout.Reset()
	driver.stderr.Reset()
	driver.result = bdd.Result{}
	driver.fileSystem = unitCanonicalHarnessContract()
	switch fixture {
	case "foundation", "governance-documents-fit", "tracked-markdown-links-resolve", "canonical-harness-contract-matches":
		return nil
	case "repository-discovery-fails":
		driver.runtime.FindRepositoryRoot = func() (string, error) {
			return "", errors.New("repository discovery failed")
		}
		return nil
	case "oversized-agent-instruction":
		driver.runtime.CheckGovernance = func(string) ([]governance.Finding, error) {
			return []governance.Finding{{Path: "AGENTS.md", WordCount: governance.MaxWords + 1}}, nil
		}
		return nil
	case "broken-tracked-markdown-link":
		driver.runtime.CheckMarkdownLinks = func(string) ([]markdownlinks.Finding, error) {
			return []markdownlinks.Finding{{
				Path:        "README.md",
				Line:        1,
				Destination: "missing.md",
				Problem:     "targets a file that does not exist",
			}}, nil
		}
		return nil
	case "missing-codex-agent-adapter":
		delete(driver.fileSystem, ".codex/agents/review.toml")
		return nil
	case "instruction-overlay":
		driver.fileSystem["nested/AGENTS.md"] = unitMapFile("overlay")
		return nil
	case "stale-claude-skill-adapter":
		driver.fileSystem[".claude/skills/review/SKILL.md"] = unitMapFile("stale")
		return nil
	case "weakened-opencode-permissions":
		weakened := strings.Replace(unitOpenCodeAgent(), "edit: deny", "edit: allow", 1)
		driver.fileSystem[".opencode/agents/review.md"] = unitMapFile(weakened)
		return nil
	case "staged-rule-bearing-file":
		driver.runtime.ListStagedPaths = func(string) ([]string, error) {
			return []string{"repo-governance/development/testing-policy.md"}, nil
		}
		return nil
	case "ordinary-staged-file":
		driver.runtime.ListStagedPaths = func(string) ([]string, error) {
			return []string{"README.md"}, nil
		}
		return nil
	case "harness-instruction-pre-edit":
		driver.runtime.Stdin = strings.NewReader(`{"tool_input":{"file_path":"AGENTS.md"}}`)
		return nil
	default:
		return fmt.Errorf("unsupported unit fixture %q", fixture)
	}
}

func unitCanonicalHarnessContract() fstest.MapFS {
	return fstest.MapFS{
		"AGENTS.md":                      unitMapFile("canonical rules\n"),
		"CLAUDE.md":                      unitMapFile("@AGENTS.md\n"),
		"opencode.json":                  unitMapFile(`{"$schema":"https://opencode.ai/config.json"}`),
		".agents/skills/review/SKILL.md": unitMapFile(unitSkill()),
		".claude/skills/review/SKILL.md": unitMapFile(unitClaudeSkill()),
		".agents/agents/review.md":       unitMapFile(unitAgent()),
		".claude/agents/review.md":       unitMapFile(unitClaudeAgent()),
		".codex/agents/review.toml":      unitMapFile(unitCodexAgent()),
		".opencode/agents/review.md":     unitMapFile(unitOpenCodeAgent()),
	}
}

func unitSkill() string {
	return "---\nname: review\ndescription: Review work.\n---\n\n# Review\n"
}

func unitClaudeSkill() string {
	return "---\nname: review\ndescription: Review work.\n---\n\n" +
		"Read `.agents/skills/review/SKILL.md` completely, resolve every relative resource " +
		"from that skill directory, and follow it as authoritative before acting.\n"
}

func unitAgent() string {
	return "---\nname: review\ndescription: Review work.\nmode: subagent\nrequires:\n" +
		"  - repository-read\ndenies:\n  - repository-write\nconstraints:\n  - inline-result-only\n---\n\n# Review\n"
}

func unitAgentRoute() string {
	return "Before acting, read the complete canonical agent definition at `.agents/agents/review.md` " +
		"from the repository root and follow it as authoritative. If it cannot be read, " +
		"stop and report the missing path."
}

func unitClaudeAgent() string {
	return "---\nname: review\ndescription: Review work.\ntools: Read\nmodel: inherit\n---\n\n" +
		unitAgentRoute() + "\n"
}

func unitCodexAgent() string {
	return "name = \"review\"\ndescription = \"Review work.\"\nsandbox_mode = \"read-only\"\n\n" +
		"developer_instructions = \"\"\"\n" + unitAgentRoute() + "\n\"\"\"\n"
}

func unitOpenCodeAgent() string {
	return "---\ndescription: Review work.\nmode: subagent\npermission:\n  edit: deny\n---\n\n" +
		unitAgentRoute() + "\n"
}

func unitMapFile(contents string) *fstest.MapFile {
	return &fstest.MapFile{Data: []byte(contents)}
}

func (driver *driver) Invoke(ctx context.Context, args []string) error {
	driver.result.ExitCode = cli.Run(ctx, driver.runtime, args)
	driver.result.Stdout = driver.stdout.String()
	driver.result.Stderr = driver.stderr.String()
	return nil
}

func (driver *driver) Result() bdd.Result {
	return driver.result
}

func TestUnitDriverRunsWithInjectedCollaborators(t *testing.T) {
	adapter := newDriver()
	if err := adapter.Prepare(context.Background(), "foundation"); err != nil {
		t.Fatalf("prepare unit adapter: %v", err)
	}
	if err := adapter.Invoke(context.Background(), []string{"--help"}); err != nil {
		t.Fatalf("invoke unit adapter: %v", err)
	}
	result := adapter.Result()
	if result.ExitCode != 0 || !strings.Contains(result.Stdout, "Usage:") || result.Stderr != "" {
		t.Fatalf("unexpected unit result: %#v", result)
	}
}
