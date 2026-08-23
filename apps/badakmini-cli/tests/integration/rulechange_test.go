package integration_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/rulechange"
)

func TestIntegrationRuleChangeReadsRealStagedPaths(t *testing.T) {
	root := t.TempDir()
	runRuleChangeGit(t, root, "init", "--quiet")
	paths := []string{"ordinary.md", "space name.md", "line\nbreak.md"}
	for _, path := range paths {
		writeRuleChangeFile(t, root, path, "content")
	}
	slices.Sort(paths)
	runRuleChangeGit(t, root, append([]string{"add", "--"}, paths...)...)

	staged := rulechange.ParseStagedPaths(stagedRuleChangeOutput(t, root))
	if strings.Join(staged, "|") != strings.Join(paths, "|") {
		t.Fatalf("expected %q, got %q", paths, staged)
	}
}

func runRuleChangeGit(t *testing.T, root string, arguments ...string) {
	t.Helper()
	commandArguments := append([]string{"-C", root}, arguments...)
	// #nosec G204 -- the test owns the repository and every Git argument.
	command := exec.Command("git", commandArguments...)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %s: %v", arguments, output, err)
	}
}

func writeRuleChangeFile(t *testing.T, root, relativePath, contents string) {
	t.Helper()
	filePath := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(filePath), 0o750); err != nil {
		t.Fatalf("create parent for %q: %v", relativePath, err)
	}
	if err := os.WriteFile(filePath, []byte(contents), 0o600); err != nil {
		t.Fatalf("write %q: %v", relativePath, err)
	}
}

func stagedRuleChangeOutput(t *testing.T, root string) []byte {
	t.Helper()
	// #nosec G204 -- the integration test owns the repository root.
	output, err := exec.Command("git", "-C", root, "diff", "--cached", "--name-only", "-z").Output()
	if err != nil {
		t.Fatalf("list staged paths: %v", err)
	}
	return output
}
