package integration_test

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
)

func TestIntegrationCLIApplicationBoundaryFailures(t *testing.T) {
	t.Run("usage and discovery", testUsageAndDiscoveryFailures)
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
	if exitCode := runRuntime(t, []string{"harness"}, func(*cli.Runtime) {}); exitCode != 0 {
		t.Fatalf("expected group help success, got %d", exitCode)
	}
	retired := []string{"harness", "instruction-size", "validate"}
	if exitCode := runRuntime(t, retired, func(*cli.Runtime) {}); exitCode != 2 {
		t.Fatalf("expected a retired check name to be an invalid invocation, got %d", exitCode)
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
	if exitCode := runRuntime(t, []string{"harness", "rule-change", "validate"}, func(runtime *cli.Runtime) {
		runtime.FindRepositoryRoot = func() (string, error) { return "", errors.New("no repository") }
	}); exitCode != 1 {
		t.Fatalf("expected repository failure, got %d", exitCode)
	}
	if exitCode := runRuntime(t, []string{"harness", "rule-change", "validate"}, func(runtime *cli.Runtime) {
		runtime.Stderr = integrationWriteFailure{}
		runtime.FindRepositoryRoot = func() (string, error) { return "", errors.New("no repository") }
	}); exitCode != 1 {
		t.Fatalf("expected repository diagnostic failure, got %d", exitCode)
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
		ListStagedPaths:    func(string) ([]string, error) { return nil, nil },
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
