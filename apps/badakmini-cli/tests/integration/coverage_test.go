package integration_test

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
)

func TestIntegrationCLIApplicationBoundaryFailures(t *testing.T) {
	t.Run("usage and discovery", testUsageAndDiscoveryFailures)
	t.Run("instruction size", testInstructionSizeFailures)
	t.Run("Markdown links", testMarkdownLinkFailures)
	t.Run("capability parity", testCapabilityParityFailures)
	t.Run("staged rule change", testStagedRuleChangeFailures)
	t.Run("hook rule change", testHookRuleChangeFailures)
}

func testUsageAndDiscoveryFailures(t *testing.T) {
	if exitCode := runRuntime(t, []string{"-h"}, func(*cli.Runtime) {}); exitCode != 0 {
		t.Fatalf("expected short help success, got %d", exitCode)
	}
	if exitCode := runRuntime(t, []string{"not", "a", "command"}, func(*cli.Runtime) {}); exitCode != 2 {
		t.Fatalf("expected invalid invocation, got %d", exitCode)
	}
	if exitCode := runRuntime(t, []string{"--help"}, func(runtime *cli.Runtime) {
		runtime.Stdout = integrationWriteFailure{}
	}); exitCode != 1 {
		t.Fatalf("expected help write failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, []string{"not", "a", "command"}, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
	}); exitCode != 1 {
		t.Fatalf("expected usage write failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, []string{"harness", "markdown-links", "validate"}, func(runtime *cli.Runtime) {
		runtime.FindRepositoryRoot = func() (string, error) { return "", errors.New("no repository") }
	}); exitCode != 1 {
		t.Fatalf("expected repository failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, []string{"harness", "markdown-links", "validate"}, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.FindRepositoryRoot = func() (string, error) { return "", errors.New("no repository") }
	}); exitCode != 1 {
		t.Fatalf("expected repository diagnostic failure, got %d", exitCode)
	}
}

func testInstructionSizeFailures(t *testing.T) {
	args := []string{"harness", "instruction-size", "validate"}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.CheckGovernance = func(string) ([]governance.Finding, error) { return nil, errors.New("check failed") }
	}); exitCode != 1 {
		t.Fatalf("expected governance error, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.CheckGovernance = func(string) ([]governance.Finding, error) { return nil, errors.New("check failed") }
	}); exitCode != 1 {
		t.Fatalf("expected governance diagnostic failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stdout = integrationWriteFailure{}
	}); exitCode != 1 {
		t.Fatalf("expected governance success write failure, got %d", exitCode)
	}
	findings := []governance.Finding{{Path: "AGENTS.md", WordCount: 501}, {Path: "CLAUDE.md", WordCount: 502}}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.CheckGovernance = func(string) ([]governance.Finding, error) { return findings, nil }
	}); exitCode != 1 {
		t.Fatalf("expected governance findings, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = &integrationFailAfterWriter{succeed: len(findings)}
		runtime.CheckGovernance = func(string) ([]governance.Finding, error) { return findings, nil }
	}); exitCode != 1 {
		t.Fatalf("expected governance guidance write failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.CheckGovernance = func(string) ([]governance.Finding, error) { return findings, nil }
	}); exitCode != 1 {
		t.Fatalf("expected governance finding write failure, got %d", exitCode)
	}
}

func testMarkdownLinkFailures(t *testing.T) {
	args := []string{"harness", "markdown-links", "validate"}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.CheckMarkdownLinks = func(string) ([]markdownlinks.Finding, error) { return nil, errors.New("check failed") }
	}); exitCode != 1 {
		t.Fatalf("expected Markdown error, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.CheckMarkdownLinks = func(string) ([]markdownlinks.Finding, error) { return nil, errors.New("check failed") }
	}); exitCode != 1 {
		t.Fatalf("expected Markdown diagnostic failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stdout = integrationWriteFailure{}
	}); exitCode != 1 {
		t.Fatalf("expected Markdown success write failure, got %d", exitCode)
	}
	findings := []markdownlinks.Finding{{Path: "README.md", Line: 2, Destination: "missing.md", Problem: "does not exist"}}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.CheckMarkdownLinks = func(string) ([]markdownlinks.Finding, error) { return findings, nil }
	}); exitCode != 1 {
		t.Fatalf("expected Markdown findings, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.CheckMarkdownLinks = func(string) ([]markdownlinks.Finding, error) { return findings, nil }
	}); exitCode != 1 {
		t.Fatalf("expected Markdown diagnostic write failure, got %d", exitCode)
	}
}

func testCapabilityParityFailures(t *testing.T) {
	args := []string{"harness", "capability-parity", "validate"}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.CheckParity = func(string) (parity.Report, error) { return parity.Report{}, errors.New("check failed") }
	}); exitCode != 1 {
		t.Fatalf("expected parity error, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.CheckParity = func(string) (parity.Report, error) { return parity.Report{}, errors.New("check failed") }
	}); exitCode != 1 {
		t.Fatalf("expected parity diagnostic failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stdout = integrationWriteFailure{}
	}); exitCode != 1 {
		t.Fatalf("expected parity success write failure, got %d", exitCode)
	}
	findings := []parity.Finding{{Capability: "skill", Harness: "Codex", Missing: []string{"review"}}}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.CheckParity = func(string) (parity.Report, error) { return parity.Report{Findings: findings}, nil }
	}); exitCode != 1 {
		t.Fatalf("expected parity findings, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = &integrationFailAfterWriter{succeed: len(findings)}
		runtime.CheckParity = func(string) (parity.Report, error) { return parity.Report{Findings: findings}, nil }
	}); exitCode != 1 {
		t.Fatalf("expected parity policy write failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.CheckParity = func(string) (parity.Report, error) { return parity.Report{Findings: findings}, nil }
	}); exitCode != 1 {
		t.Fatalf("expected parity finding write failure, got %d", exitCode)
	}
}

func testStagedRuleChangeFailures(t *testing.T) {
	args := []string{"harness", "rule-change", "validate"}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.ListStagedPaths = func(string) ([]string, error) { return nil, errors.New("Git failed") }
	}); exitCode != 1 {
		t.Fatalf("expected staged path error, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.ListStagedPaths = func(string) ([]string, error) { return nil, errors.New("Git failed") }
	}); exitCode != 1 {
		t.Fatalf("expected staged path diagnostic failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.ListStagedPaths = func(string) ([]string, error) { return []string{"AGENTS.md"}, nil }
		runtime.Stdout = integrationWriteFailure{}
	}); exitCode != 1 {
		t.Fatalf("expected staged notice write failure, got %d", exitCode)
	}
}

func testHookRuleChangeFailures(t *testing.T) {
	args := []string{"harness", "rule-change", "hook"}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stdin = strings.NewReader(`{"tool_input":{"file_path":"README.md"}}`)
	}); exitCode != 0 {
		t.Fatalf("expected ordinary hook silence, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stdin = integrationReadFailure{}
	}); exitCode != 0 {
		t.Fatalf("expected nonblocking hook read failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, args, func(runtime *cli.Runtime) {
		runtime.Stdin = strings.NewReader(`{"tool_input":{"file_path":"AGENTS.md"}}`)
		runtime.Stdout = integrationWriteFailure{}
	}); exitCode != 0 {
		t.Fatalf("expected nonblocking hook write failure, got %d", exitCode)
	}
}

func runRuntime(t *testing.T, args []string, adjust func(*cli.Runtime)) int {
	t.Helper()
	var stdout, stderr bytes.Buffer
	runtime := cli.Runtime{
		Stdin:              strings.NewReader(""),
		Stdout:             &stdout,
		Stderr:             &stderr,
		FindRepositoryRoot: func() (string, error) { return "repository", nil },
		CheckGovernance:    func(string) ([]governance.Finding, error) { return nil, nil },
		CheckMarkdownLinks: func(string) ([]markdownlinks.Finding, error) { return nil, nil },
		ListStagedPaths:    func(string) ([]string, error) { return nil, nil },
		CheckParity:        func(string) (parity.Report, error) { return parity.Report{}, nil },
	}
	adjust(&runtime)
	return cli.Run(context.Background(), runtime, args)
}

type integrationWriteFailure struct{}

func (integrationWriteFailure) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

type integrationReadFailure struct{}

func (integrationReadFailure) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

type integrationFailAfterWriter struct {
	succeed int
	writes  int
}

func (writer *integrationFailAfterWriter) Write(message []byte) (int, error) {
	if writer.writes >= writer.succeed {
		return 0, errors.New("write failed")
	}
	writer.writes++
	return len(message), nil
}
