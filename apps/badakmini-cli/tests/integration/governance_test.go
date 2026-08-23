package integration_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
)

func TestIntegrationGovernanceChecksRealRepositoryFilesystem(t *testing.T) {
	root := t.TempDir()
	writeGovernanceFile(t, root, "AGENTS.md", strings.Repeat("word ", governance.MaxWords+1))
	writeGovernanceFile(t, root, "CLAUDE.md", "short")
	writeGovernanceFile(t, root, "repo-governance/policy.md", "short")

	findings, err := governance.CheckFS(os.DirFS(root))
	if err != nil {
		t.Fatalf("check real repository fixture: %v", err)
	}
	if len(findings) != 1 || findings[0].Path != "AGENTS.md" {
		t.Fatalf("expected the oversized real file to be reported, got %#v", findings)
	}
}

func TestIntegrationGovernanceReportsRealSymlinkFailure(t *testing.T) {
	root := t.TempDir()
	writeGovernanceFile(t, root, "AGENTS.md", "short")
	writeGovernanceFile(t, root, "CLAUDE.md", "short")
	if err := os.Mkdir(filepath.Join(root, "repo-governance"), 0o750); err != nil {
		t.Fatalf("create governance directory: %v", err)
	}
	harnessPath := filepath.Join(root, ".claude")
	if err := os.Symlink(harnessPath, harnessPath); err != nil {
		t.Fatalf("create cyclic harness symlink: %v", err)
	}

	_, err := governance.CheckFS(os.DirFS(root))
	if err == nil || !strings.Contains(err.Error(), "inspect .claude") {
		t.Fatalf("expected contextualized symlink failure, got %v", err)
	}
}

func writeGovernanceFile(t *testing.T, root, relativePath, contents string) {
	t.Helper()
	filePath := filepath.Join(root, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(filePath), 0o750); err != nil {
		t.Fatalf("create parent for %s: %v", relativePath, err)
	}
	if err := os.WriteFile(filePath, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %s: %v", relativePath, err)
	}
}
