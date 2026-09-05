package integration_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
)

const (
	canonicalSkillFixture = `---
name: review
description: Review work.
---

# Review

Canonical skill.
`
	canonicalSkillRoute = "Read `.agents/skills/review/SKILL.md` completely, resolve every relative resource " +
		"from that skill directory, and follow it as authoritative before acting."
	claudeSkillFixture = `---
name: review
description: Review work.
---

` + canonicalSkillRoute + "\n"
	canonicalAgentFixture = `---
name: review
description: Review work.
mode: subagent
requires:
  - repository-read
denies:
  - repository-write
constraints:
  - inline-result-only
---

# Review

Canonical agent.
`
	canonicalAgentRoute = "Before acting, read the complete canonical agent definition at `.agents/agents/review.md` " +
		"from the repository root and follow it as authoritative. If it cannot be read, " +
		"stop and report the missing path."
	claudeAgentFixture = `---
name: review
description: Review work.
tools: Read
model: inherit
---

` + canonicalAgentRoute + "\n"
	codexAgentFixture = `name = "review"
description = "Review work."
sandbox_mode = "read-only"

developer_instructions = """
` + canonicalAgentRoute + `
"""
`
	opencodeAgentFixture = `---
description: Review work.
mode: subagent
permission:
  edit: deny
---

` + canonicalAgentRoute + "\n"
)

func writeCanonicalHarnessFixture(t *testing.T, root string) {
	t.Helper()
	for path, contents := range map[string]string{
		"AGENTS.md":                      "canonical rules\n",
		"CLAUDE.md":                      "@AGENTS.md\n",
		"opencode.json":                  "{\"$schema\":\"https://opencode.ai/config.json\"}\n",
		".agents/skills/review/SKILL.md": canonicalSkillFixture,
		".claude/skills/review/SKILL.md": claudeSkillFixture,
		".agents/agents/review.md":       canonicalAgentFixture,
		".claude/agents/review.md":       claudeAgentFixture,
		".codex/agents/review.toml":      codexAgentFixture,
		".opencode/agents/review.md":     opencodeAgentFixture,
	} {
		writeParityFile(t, root, path, contents)
	}
}

func TestIntegrationParityReadsRealHarnessDirectories(t *testing.T) {
	root := t.TempDir()
	writeCanonicalHarnessFixture(t, root)

	report, err := parity.CheckFS(os.DirFS(root))
	if err != nil {
		t.Fatalf("check real harness fixture: %v", err)
	}
	if len(report.Findings) != 0 {
		t.Fatalf("expected matching real harnesses, got %#v", report.Findings)
	}
}

func TestIntegrationParityReportsRealHarnessSymlinkFailure(t *testing.T) {
	root := t.TempDir()
	writeParityFile(t, root, "AGENTS.md", "canonical rules\n")
	writeParityFile(t, root, "CLAUDE.md", "@AGENTS.md\n")
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
