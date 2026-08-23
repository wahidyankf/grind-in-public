package unit_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

// driver executes command orchestration with injected collaborators only. The
// canonical scenarios added in Phase 2 will teach Prepare which deterministic
// fixture each Given step requests.
type driver struct {
	runtime cli.Runtime
	stdout  bytes.Buffer
	stderr  bytes.Buffer
	result  bdd.Result
}

func newDriver() *driver {
	driver := &driver{}
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
		CheckParity:        func(string) ([]parity.Finding, error) { return nil, nil },
	}
	return driver
}

func (driver *driver) Prepare(_ context.Context, fixture string) error {
	driver.stdout.Reset()
	driver.stderr.Reset()
	driver.result = bdd.Result{}
	switch fixture {
	case "foundation", "governance-documents-fit", "tracked-markdown-links-resolve", "harness-capabilities-match":
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
	case "harness-missing-shared-subagent":
		driver.runtime.CheckParity = func(string) ([]parity.Finding, error) {
			return []parity.Finding{{
				Capability: "subagent",
				Harness:    "Codex",
				Missing:    []string{"review"},
			}}, nil
		}
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
