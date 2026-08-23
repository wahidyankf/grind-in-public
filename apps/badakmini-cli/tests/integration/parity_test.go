package integration_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
)

func TestIntegrationParityReadsRealHarnessDirectories(t *testing.T) {
	root := t.TempDir()
	for _, path := range []string{
		".claude/agents/drill-reviewer.md",
		".codex/agents/drill-reviewer.toml",
		".opencode/agents/drill-reviewer.md",
		".claude/skills/review/SKILL.md",
		".agents/skills/review/SKILL.md",
	} {
		writeParityFile(t, root, path, "fixture")
	}

	findings, err := parity.CheckFS(os.DirFS(root))
	if err != nil {
		t.Fatalf("check real harness fixture: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected matching real harnesses, got %#v", findings)
	}
}

func TestIntegrationParityReportsRealHarnessSymlinkFailure(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, ".claude"), 0o750); err != nil {
		t.Fatalf("create Claude harness directory: %v", err)
	}
	agentsPath := filepath.Join(root, ".claude", "agents")
	if err := os.Symlink(agentsPath, agentsPath); err != nil {
		t.Fatalf("create cyclic agents symlink: %v", err)
	}

	_, err := parity.CheckFS(os.DirFS(root))
	if err == nil || !strings.Contains(err.Error(), "read .claude/agents") {
		t.Fatalf("expected contextualized harness directory read error, got %v", err)
	}
}

func writeParityFile(t *testing.T, root, relativePath, contents string) {
	t.Helper()
	filePath := filepath.Join(root, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(filePath), 0o750); err != nil {
		t.Fatalf("create parent for %s: %v", relativePath, err)
	}
	if err := os.WriteFile(filePath, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %s: %v", relativePath, err)
	}
}
