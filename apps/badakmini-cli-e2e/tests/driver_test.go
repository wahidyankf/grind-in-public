package e2e_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const processTimeout = 30 * time.Second

type processDriver struct {
	testing  *testing.T
	binary   string
	root     string
	setupErr error
	stdin    string
	result   commandResult
}

type commandResult struct {
	ExitCode int
	Stdout   string
	Stderr   string
}

func newProcessDriver(t *testing.T) *processDriver {
	t.Helper()
	binary, err := resolveBinary(os.Getenv("BADAKMINI_BIN"))
	return &processDriver{testing: t, binary: binary, root: t.TempDir(), setupErr: err}
}

func resolveBinary(candidate string) (string, error) {
	if candidate == "" {
		return "", errors.New("BADAKMINI_BIN must name an absolute executable")
	}
	if !filepath.IsAbs(candidate) {
		return "", fmt.Errorf("BADAKMINI_BIN must be absolute: %q", candidate)
	}
	info, err := os.Stat(candidate) // #nosec G703 -- E2E contract validates the caller-provided executable path.
	if err != nil {
		return "", fmt.Errorf("inspect BADAKMINI_BIN %q: %w", candidate, err)
	}
	if !info.Mode().IsRegular() || info.Mode()&0o111 == 0 {
		return "", fmt.Errorf("BADAKMINI_BIN must be an executable file: %q", candidate)
	}
	return candidate, nil
}

//nolint:cyclop // One switch keeps the canonical fixture names visible and exhaustive.
func (driver *processDriver) Prepare(_ context.Context, fixture string) error {
	if driver.setupErr != nil {
		return driver.setupErr
	}
	driver.stdin = ""
	driver.result = commandResult{}
	if fixture != "repository-discovery-fails" {
		driver.initializeGit()
	}
	switch fixture {
	case "repository-discovery-fails":
		return nil
	case "governance-documents-fit":
		driver.writeGovernance("short")
	case "oversized-agent-instruction":
		driver.writeGovernance(strings.Repeat("word ", 501))
	case "tracked-markdown-links-resolve":
		driver.writeFile("README.md", "[Guide](docs/guide.md)\n")
		driver.writeFile("docs/guide.md", "# Guide\n")
		driver.runGit("add", "README.md", "docs/guide.md")
	case "broken-tracked-markdown-link":
		driver.writeFile("README.md", "[Missing](missing.md)\n")
		driver.runGit("add", "README.md")
	case "harness-capabilities-match":
		for _, path := range []string{
			".claude/agents/review.md",
			".codex/agents/review.toml",
			".opencode/agents/review.md",
		} {
			driver.writeFile(path, "fixture")
		}
	case "harness-missing-shared-subagent":
		driver.writeFile(".claude/agents/review.md", "fixture")
	case "staged-rule-bearing-file":
		driver.stageFile("repo-governance/development/testing-policy.md")
	case "ordinary-staged-file":
		driver.stageFile("README.md")
	case "harness-instruction-pre-edit":
		driver.stdin = `{"tool_input":{"file_path":"AGENTS.md"}}`
	default:
		return fmt.Errorf("unsupported E2E fixture: %s", fixture)
	}
	return nil
}

func (driver *processDriver) Invoke(parent context.Context, arguments []string) error {
	if driver.setupErr != nil {
		return driver.setupErr
	}
	ctx, cancel := context.WithTimeout(parent, processTimeout)
	defer cancel()

	// #nosec G204 -- resolveBinary permits only the configured absolute executable.
	command := exec.CommandContext(ctx, driver.binary, arguments...)
	command.Dir = driver.root
	command.Stdin = strings.NewReader(driver.stdin)
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	driver.result = commandResult{Stdout: stdout.String(), Stderr: stderr.String()}
	if ctx.Err() != nil {
		return fmt.Errorf("Badak Mini exceeded %s: %w", processTimeout, ctx.Err())
	}
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		driver.result.ExitCode = exitError.ExitCode()
		return nil
	}
	if err != nil {
		return fmt.Errorf("run Badak Mini: %w", err)
	}
	return nil
}

func (driver *processDriver) Result() commandResult {
	return driver.result
}

func (driver *processDriver) initializeGit() {
	driver.runGit("init", "--quiet")
}

func (driver *processDriver) stageFile(path string) {
	driver.writeFile(path, "fixture")
	driver.runGit("add", "--", path)
}

func (driver *processDriver) writeGovernance(agents string) {
	driver.writeFile("AGENTS.md", agents)
	driver.writeFile("CLAUDE.md", "short")
	driver.writeFile("repo-governance/policy.md", "short")
}

func (driver *processDriver) writeFile(path, content string) {
	driver.testing.Helper()
	fullPath := filepath.Join(driver.root, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o750); err != nil {
		driver.testing.Fatalf("create fixture parent for %s: %v", path, err)
	}
	if err := os.WriteFile(fullPath, []byte(content), 0o600); err != nil {
		driver.testing.Fatalf("write fixture %s: %v", path, err)
	}
}

func (driver *processDriver) runGit(arguments ...string) {
	driver.testing.Helper()
	// #nosec G204 -- fixture Git arguments are test-owned constants.
	command := exec.Command("git", append([]string{"-C", driver.root}, arguments...)...)
	if output, err := command.CombinedOutput(); err != nil {
		driver.testing.Fatalf("git %s: %v\n%s", strings.Join(arguments, " "), err, output)
	}
}

func TestProcessDriverBinaryContract(t *testing.T) {
	fixture := filepath.Join(t.TempDir(), "badak-mini-fixture")
	// #nosec G306 -- executable fixture requires its owner execute bit.
	if err := os.WriteFile(fixture, []byte("#!/bin/sh\nprintf '%s' \"$*\"\n"), 0o700); err != nil {
		t.Fatalf("write executable fixture: %v", err)
	}

	for _, candidate := range []string{"", "badak-mini", filepath.Join(t.TempDir(), "missing")} {
		if _, err := resolveBinary(candidate); err == nil {
			t.Errorf("resolveBinary(%q) succeeded", candidate)
		}
	}
	nonExecutable := filepath.Join(t.TempDir(), "not-executable")
	if err := os.WriteFile(nonExecutable, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("write non-executable fixture: %v", err)
	}
	if _, err := resolveBinary(nonExecutable); err == nil {
		t.Error("resolveBinary accepted a non-executable file")
	}

	t.Setenv("BADAKMINI_BIN", fixture)
	driver := newProcessDriver(t)
	if err := driver.Prepare(context.Background(), "repository-discovery-fails"); err != nil {
		t.Fatalf("prepare process driver: %v", err)
	}
	if err := driver.Invoke(context.Background(), []string{"only-absolute"}); err != nil {
		t.Fatalf("invoke process driver: %v", err)
	}
	if result := driver.Result(); result.ExitCode != 0 || result.Stdout != "only-absolute" {
		t.Fatalf("unexpected process result: %#v", result)
	}
}
