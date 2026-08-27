package integration_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/rulechange"
)

func TestIntegrationCLIExecutesRealValidationAdapters(t *testing.T) {
	for _, test := range cliAdapterCases() {
		t.Run(test.name, func(t *testing.T) {
			exitCode, stdout, stderr := runCLIAtRoot(test.args, test.prepare(t))
			if exitCode != 0 || !strings.Contains(stdout, test.expectText) {
				t.Fatalf("expected successful adapter run, got exit %d, stdout %q, stderr %q", exitCode, stdout, stderr)
			}
		})
	}
}

type cliAdapterCase struct {
	name       string
	args       []string
	prepare    func(*testing.T) string
	expectText string
}

func cliAdapterCases() []cliAdapterCase {
	return []cliAdapterCase{
		{
			name: "instruction size",
			args: []string{"harness", "instruction-size", "validate"},
			prepare: func(t *testing.T) string {
				t.Helper()
				root := t.TempDir()
				writeGovernanceFile(t, root, "AGENTS.md", "short")
				writeGovernanceFile(t, root, "CLAUDE.md", "short")
				writeGovernanceFile(t, root, "repo-governance/policy.md", "short")
				return root
			},
			expectText: "Governance word counts",
		},
		{
			name: "Markdown links",
			args: []string{"harness", "markdown-links", "validate"},
			prepare: func(t *testing.T) string {
				t.Helper()
				root := newMarkdownRepository(t)
				writeMarkdownFile(t, root, "README.md", "# Repository\n")
				runMarkdownGit(t, root, "add", "README.md")
				return root
			},
			expectText: "Markdown links are valid",
		},
		{
			name: "staged rule change",
			args: []string{"harness", "rule-change", "validate"},
			prepare: func(t *testing.T) string {
				t.Helper()
				root := t.TempDir()
				runRuleChangeGit(t, root, "init", "--quiet")
				writeRuleChangeFile(t, root, "AGENTS.md", "short")
				runRuleChangeGit(t, root, "add", "AGENTS.md")
				return root
			},
			expectText: "Rules Propagation automatically triggered",
		},
		{
			name: "capability parity",
			args: []string{"harness", "capability-parity", "validate"},
			prepare: func(t *testing.T) string {
				t.Helper()
				return t.TempDir()
			},
			expectText: "Every harness exposes",
		},
	}
}

func TestIntegrationCLIReportsRealValidationFinding(t *testing.T) {
	root := t.TempDir()
	writeGovernanceFile(t, root, "AGENTS.md", strings.Repeat("word ", governance.MaxWords+1))
	writeGovernanceFile(t, root, "CLAUDE.md", "short")
	writeGovernanceFile(t, root, "repo-governance/policy.md", "short")

	exitCode, _, stderr := runCLIAtRoot([]string{"harness", "instruction-size", "validate"}, root)
	if exitCode != 1 || !strings.Contains(stderr, fmt.Sprintf("contains %d words", governance.MaxWords+1)) {
		t.Fatalf("expected validation diagnostic, got exit %d and stderr %q", exitCode, stderr)
	}
}

func TestIntegrationCLIReportsRealGitFailure(t *testing.T) {
	exitCode, _, stderr := runCLIAtRoot([]string{"harness", "rule-change", "validate"}, t.TempDir())
	if exitCode != 1 || !strings.Contains(stderr, "list staged paths") {
		t.Fatalf("expected Git diagnostic, got exit %d and stderr %q", exitCode, stderr)
	}
}

func runCLIAtRoot(args []string, root string) (int, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runtime := cli.Runtime{
		Stdin:              strings.NewReader(""),
		Stdout:             &stdout,
		Stderr:             &stderr,
		FindRepositoryRoot: func() (string, error) { return root, nil },
		CheckGovernance:    func(root string) ([]governance.Finding, error) { return governance.CheckFS(os.DirFS(root)) },
		CheckMarkdownLinks: func(root string) ([]markdownlinks.Finding, error) {
			return markdownlinks.Check(root, realMarkdownRuntime())
		},
		ListStagedPaths: realStagedPaths,
		CheckParity:     func(root string) ([]parity.Finding, error) { return parity.CheckFS(os.DirFS(root)) },
	}
	exitCode := cli.Run(context.Background(), runtime, args)
	return exitCode, stdout.String(), stderr.String()
}

func realStagedPaths(root string) ([]string, error) {
	// #nosec G204 -- the integration test owns the repository root.
	output, err := exec.Command("git", "-C", root, "diff", "--cached", "--name-only", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("list staged paths: %w", err)
	}
	return rulechange.ParseStagedPaths(output), nil
}
