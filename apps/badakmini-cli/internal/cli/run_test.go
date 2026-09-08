package cli

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

const (
	ruleChangeValidateCommand = "harness rule-change validate"
	ruleChangeHookCommand     = "harness rule-change hook"
)

var supportedCommands = []string{
	ruleChangeValidateCommand,
	ruleChangeHookCommand,
}

func TestRunPrintsHelpWithoutRepositoryDiscovery(t *testing.T) {
	runtime, stdout, stderr := runtimeDouble(t)
	runtime.FindRepositoryRoot = func() (string, error) {
		t.Fatal("help must not discover the repository")
		return "", nil
	}

	exitCode := Run(context.Background(), runtime, []string{"--help"})
	if exitCode != 0 || stderr.Len() != 0 {
		t.Fatalf("expected successful help, got exit %d and stderr %q", exitCode, stderr.String())
	}
	for _, command := range supportedCommands {
		if !strings.Contains(stdout.String(), command) {
			t.Fatalf("expected %q in usage, got %q", command, stdout.String())
		}
	}
}

func TestRunRejectsUnsupportedCommandsBeforeRepositoryDiscovery(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "unknown command", args: []string{"unknown"}},
		{name: "unknown flag", args: []string{"--unknown"}},
		{
			name: "excess leaf argument",
			args: []string{"harness", "rule-change", "validate", "unexpected"},
		},
		{
			// A name RHINO now owns. A stale hook still calling it must fail
			// rather than report a clean run of a check that no longer exists.
			name: "retired check name",
			args: []string{"harness", "instruction-size", "validate"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runtime, _, stderr := runtimeDouble(t)
			runtime.FindRepositoryRoot = func() (string, error) {
				t.Fatal("invalid usage must not discover the repository")
				return "", nil
			}

			exitCode := Run(context.Background(), runtime, test.args)
			if exitCode != invalidInvocationExitCode || !strings.Contains(stderr.String(), "Usage:") {
				t.Fatalf("expected usage failure, got exit %d and stderr %q", exitCode, stderr.String())
			}
		})
	}
}

func TestRunPrintsGroupHelpWhenNoLeafIsNamed(t *testing.T) {
	runtime, stdout, stderr := runtimeDouble(t)
	runtime.FindRepositoryRoot = func() (string, error) {
		t.Fatal("group help must not discover the repository")
		return "", nil
	}

	exitCode := Run(context.Background(), runtime, []string{"harness"})
	if exitCode != 0 || stderr.Len() != 0 {
		t.Fatalf("expected successful group help, got exit %d and stderr %q", exitCode, stderr.String())
	}
	if !strings.Contains(stdout.String(), "rule-change") {
		t.Fatalf("expected the group's leaves in usage, got %q", stdout.String())
	}
}

func TestRunReturnsTheExitCodeAnActionReports(t *testing.T) {
	runtime, stdout, stderr := runtimeDouble(t)
	runtime.ListStagedPaths = func(string) ([]string, error) { return nil, errors.New("Git failed") }

	exitCode := Run(context.Background(), runtime, strings.Fields(ruleChangeValidateCommand))
	if exitCode != 1 {
		t.Fatalf("expected the action's failing exit code, got %d", exitCode)
	}
	if stdout.Len() != 0 || !strings.Contains(stderr.String(), "Git failed") {
		t.Fatalf("expected the diagnostic on stderr alone, got %q and %q", stdout.String(), stderr.String())
	}
}

func TestRunReportsUsageWriteFailures(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		stdout io.Writer
		stderr io.Writer
	}{
		{name: "help", args: []string{"--help"}, stdout: writeFailure{}, stderr: io.Discard},
		{name: "invalid invocation", args: []string{"unknown"}, stdout: io.Discard, stderr: writeFailure{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runtime := Runtime{
				Stdout: test.stdout,
				Stderr: test.stderr,
				FindRepositoryRoot: func() (string, error) {
					t.Fatal("usage handling must not discover the repository")
					return "", nil
				},
			}
			if exitCode := Run(context.Background(), runtime, test.args); exitCode != 1 {
				t.Fatalf("expected output failure exit, got %d", exitCode)
			}
		})
	}
}

func TestRunReportsRepositoryDiscoveryFailure(t *testing.T) {
	runtime, _, stderr := runtimeDouble(t)
	runtime.FindRepositoryRoot = func() (string, error) {
		return "", errors.New("not a repository")
	}

	exitCode := Run(context.Background(), runtime, strings.Fields(ruleChangeValidateCommand))
	if exitCode != 1 || !strings.Contains(stderr.String(), "could not find") {
		t.Fatalf("expected repository discovery failure, got exit %d and stderr %q", exitCode, stderr.String())
	}
}

func TestRunReturnsFailureWhenDiscoveryDiagnosticCannotBeWritten(t *testing.T) {
	runtime, _, _ := runtimeDouble(t)
	runtime.Stderr = writeFailure{}
	runtime.FindRepositoryRoot = func() (string, error) {
		return "", errors.New("not a repository")
	}

	if exitCode := Run(context.Background(), runtime, strings.Fields(ruleChangeValidateCommand)); exitCode != 1 {
		t.Fatalf("expected failure exit, got %d", exitCode)
	}
}

func TestRunDispatchesEverySupportedCommandThroughTheRuntime(t *testing.T) {
	for _, command := range supportedCommands {
		t.Run(command, func(t *testing.T) { assertDispatchedCommand(t, command) })
	}
}

func assertDispatchedCommand(t *testing.T, command string) {
	t.Helper()
	runtime, stdout, stderr := runtimeDouble(t)
	called := false
	switch command {
	case ruleChangeValidateCommand:
		runtime.ListStagedPaths = func(string) ([]string, error) { called = true; return nil, nil }
	case ruleChangeHookCommand:
		runtime.Stdin = strings.NewReader(`{"tool_input":{"file_path":"AGENTS.md"}}`)
	}

	exitCode := Run(context.Background(), runtime, strings.Fields(command))
	if command == ruleChangeHookCommand {
		called = strings.Contains(stdout.String(), "PreToolUse")
	}
	if !called {
		t.Fatalf("expected %q runtime dispatch", command)
	}
	if exitCode != 0 {
		t.Fatalf("expected dispatched success, got %d", exitCode)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no orchestration stderr, got %q", stderr.String())
	}
}

func TestAnnounceHookRuleChangeHandlesRelevantAndOrdinaryPayloads(t *testing.T) {
	root := "/repository"
	relevant := `{"tool_input":{"file_path":"/repository/AGENTS.md"}}`
	var stdout bytes.Buffer

	exitCode := announceHookRuleChange(root, strings.NewReader(relevant), &stdout)
	if exitCode != 0 || !strings.Contains(stdout.String(), "PreToolUse") {
		t.Fatalf("expected hook context, got exit %d and output %q", exitCode, stdout.String())
	}

	stdout.Reset()
	exitCode = announceHookRuleChange(root, strings.NewReader(`{"tool_input":{"file_path":"README.md"}}`), &stdout)
	if exitCode != 0 || stdout.Len() != 0 {
		t.Fatalf("expected ordinary edit silence, got exit %d and output %q", exitCode, stdout.String())
	}
}

func TestAnnounceHookRuleChangeNeverBlocksOnStreamFailure(t *testing.T) {
	if exitCode := announceHookRuleChange("/repository", readFailure{}, io.Discard); exitCode != 0 {
		t.Fatalf("expected input failure to stay nonblocking, got %d", exitCode)
	}
	if exitCode := announceHookRuleChange(
		"/repository",
		strings.NewReader(`{"tool_input":{"file_path":"AGENTS.md"}}`),
		writeFailure{},
	); exitCode != 0 {
		t.Fatalf("expected output failure to stay nonblocking, got %d", exitCode)
	}
}

func TestAnnounceStagedRuleChangeCoversResultsAndStreamFailures(t *testing.T) {
	t.Run("staged path failure", func(t *testing.T) {
		runtime, _, stderr := runtimeDouble(t)
		runtime.ListStagedPaths = func(string) ([]string, error) { return nil, errors.New("git failed") }
		exitCode := announceStagedRuleChange(runtime, "repository")
		if exitCode != 1 || !strings.Contains(stderr.String(), "git failed") {
			t.Fatalf("expected staged-path failure, got exit %d and stderr %q", exitCode, stderr.String())
		}
	})

	t.Run("staged path diagnostic failure", func(t *testing.T) {
		runtime, _, _ := runtimeDouble(t)
		runtime.Stderr = writeFailure{}
		runtime.ListStagedPaths = func(string) ([]string, error) { return nil, errors.New("git failed") }
		if exitCode := announceStagedRuleChange(runtime, "repository"); exitCode != 1 {
			t.Fatalf("expected diagnostic failure exit, got %d", exitCode)
		}
	})

	t.Run("rule notice", func(t *testing.T) {
		runtime, stdout, _ := runtimeDouble(t)
		runtime.ListStagedPaths = func(string) ([]string, error) { return []string{"AGENTS.md"}, nil }
		exitCode := announceStagedRuleChange(runtime, "repository")
		if exitCode != 0 || !strings.Contains(stdout.String(), "Rules Propagation automatically triggered") {
			t.Fatalf("expected rule notice, got exit %d and stdout %q", exitCode, stdout.String())
		}
	})

	t.Run("rule notice output failure", func(t *testing.T) {
		runtime, _, _ := runtimeDouble(t)
		runtime.Stdout = writeFailure{}
		runtime.ListStagedPaths = func(string) ([]string, error) { return []string{"AGENTS.md"}, nil }
		if exitCode := announceStagedRuleChange(runtime, "repository"); exitCode != 1 {
			t.Fatalf("expected notice output failure, got %d", exitCode)
		}
	})
}

func TestWritefWrapsWriterFailure(t *testing.T) {
	err := writef(writeFailure{}, "message %d", 1)
	if err == nil || !strings.Contains(err.Error(), "write formatted output") {
		t.Fatalf("expected contextualized writer failure, got %v", err)
	}
}

type writeFailure struct{}

func (writeFailure) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

type readFailure struct{}

func (readFailure) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func runtimeDouble(t *testing.T) (Runtime, *bytes.Buffer, *bytes.Buffer) {
	t.Helper()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runtime := Runtime{
		Stdin:  strings.NewReader("payload"),
		Stdout: &stdout,
		Stderr: &stderr,
		FindRepositoryRoot: func() (string, error) {
			return "repository", nil
		},
		ListStagedPaths: func(string) ([]string, error) { return nil, nil },
	}
	return runtime, &stdout, &stderr
}
