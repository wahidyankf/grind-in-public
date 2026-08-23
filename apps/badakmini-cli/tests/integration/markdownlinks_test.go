package integration_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
)

func TestIntegrationMarkdownLinksReadsTheRealGitIndex(t *testing.T) {
	root := newMarkdownRepository(t)
	writeMarkdownFile(t, root, "README.md", "[Guide](docs/guide.md#getting-started)\n")
	writeMarkdownFile(t, root, "docs/guide.md", "# Getting Started\n")
	writeMarkdownFile(t, root, "node_modules/ignored.md", "[Broken](missing.md)\n")
	runMarkdownGit(t, root, "add", "README.md", "docs/guide.md")

	findings, err := markdownlinks.Check(root, realMarkdownRuntime())
	if err != nil {
		t.Fatalf("check real Git fixture: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected tracked links to resolve and untracked files to be ignored, got %#v", findings)
	}
}

func TestIntegrationMarkdownLinksReportsARealStagedDeletion(t *testing.T) {
	root := newMarkdownRepository(t)
	writeMarkdownFile(t, root, "README.md", "[Guide](docs/guide.md)\n")
	writeMarkdownFile(t, root, "docs/guide.md", "# Guide\n")
	runMarkdownGit(t, root, "add", "README.md", "docs/guide.md")
	runMarkdownGit(t, root, "rm", "--quiet", "--force", "docs/guide.md")

	findings, err := markdownlinks.Check(root, realMarkdownRuntime())
	if err != nil {
		t.Fatalf("check staged deletion: %v", err)
	}
	if len(findings) != 1 || findings[0].Problem != "targets a file that does not exist" {
		t.Fatalf("expected staged deletion finding, got %#v", findings)
	}
}

func TestIntegrationMarkdownLinksRejectsARealEscapingSymlink(t *testing.T) {
	root := newMarkdownRepository(t)
	externalRoot := t.TempDir()
	writeMarkdownFile(t, root, "README.md", "[External](docs/external.md)\n")
	writeMarkdownFile(t, externalRoot, "external.md", "# External\n")
	if err := os.MkdirAll(filepath.Join(root, "docs"), 0o750); err != nil {
		t.Fatalf("create docs directory: %v", err)
	}
	externalPath := filepath.Join(externalRoot, "external.md")
	linkedPath := filepath.Join(root, "docs", "external.md")
	if err := os.Symlink(externalPath, linkedPath); err != nil {
		t.Fatalf("create external symlink: %v", err)
	}
	runMarkdownGit(t, root, "add", "README.md", "docs/external.md")

	findings, err := markdownlinks.Check(root, realMarkdownRuntime())
	if err != nil {
		t.Fatalf("check symlink escape: %v", err)
	}
	if len(findings) != 1 || findings[0].Problem != "resolves outside this repository" {
		t.Fatalf("expected symlink escape finding, got %#v", findings)
	}
}

func newMarkdownRepository(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	runMarkdownGit(t, root, "init", "--quiet")
	return root
}

func runMarkdownGit(t *testing.T, root string, arguments ...string) {
	t.Helper()
	commandArguments := append([]string{"-C", root}, arguments...)
	// #nosec G204 -- the test owns the repository and every Git argument.
	command := exec.Command("git", commandArguments...)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %s: %v", arguments, output, err)
	}
}

func writeMarkdownFile(t *testing.T, root, relativePath, contents string) {
	t.Helper()
	filePath := filepath.Join(root, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(filePath), 0o750); err != nil {
		t.Fatalf("create parent for %s: %v", relativePath, err)
	}
	if err := os.WriteFile(filePath, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %s: %v", relativePath, err)
	}
}

func realMarkdownRuntime() markdownlinks.Runtime {
	return markdownlinks.Runtime{
		ReadFile:     os.ReadFile,
		Stat:         os.Stat,
		EvalSymlinks: filepath.EvalSymlinks,
		TrackedFiles: func(root string) (map[string]struct{}, error) {
			// #nosec G204 -- the integration test owns the repository root.
			output, err := exec.Command("git", "-C", root, "ls-files", "-z").Output()
			if err != nil {
				return nil, fmt.Errorf("list tracked files: %w", err)
			}
			return markdownlinks.ParseTrackedFiles(output), nil
		},
	}
}
