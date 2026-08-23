package parity

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
)

func TestCheckAcceptsHarnessesThatExposeTheSameEntries(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer", "repo-explorer")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsAHarnessMissingASubagent(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer")
	putParityFile(fileSystem, ".claude/agents/planner.md", "planner")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("expected one finding per lagging harness, got %#v", findings)
	}
	for _, finding := range findings {
		if finding.Capability != "subagent" || len(finding.Missing) != 1 || finding.Missing[0] != "planner" {
			t.Fatalf("unexpected finding: %#v", finding)
		}
		if finding.Harness == claudeHarness {
			t.Fatalf("the harness that has the agent must not be reported: %#v", finding)
		}
	}
}

func TestCheckIgnoresIndexesAndNonCapabilities(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer")
	putParityFile(fileSystem, ".claude/agents/README.md", "index")
	putParityFile(fileSystem, ".opencode/agents/README.md", "index")
	putParityFile(fileSystem, ".codex/agents/README.md", "index")
	putParityFile(fileSystem, ".claude/agents/notes.txt", "notes")
	putParityFile(fileSystem, ".claude/agents/nested/planner.md", "nested")
	putParityFile(fileSystem, ".claude/skills/reference/notes.md", "support")
	putParityFile(fileSystem, ".claude/skills/standalone.md", "not a skill")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected indexes and non-capabilities to be ignored, got %#v", findings)
	}
}

func TestCheckSkipsCapabilitiesNoHarnessUses(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	for _, finding := range findings {
		if finding.Capability == "skill" {
			t.Fatalf("expected no skill findings, got %#v", finding)
		}
	}
}

func TestCheckRequiresASkillInEveryHarnessDirectory(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer")
	putParityFile(fileSystem, ".claude/skills/review/SKILL.md", "review")
	putParityFile(fileSystem, ".claude/skills/review/reference.md", "notes")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 1 || findings[0].Harness != codexHarness || findings[0].Missing[0] != "review" {
		t.Fatalf("expected only Codex to be missing the skill, got %#v", findings)
	}
}

func TestCheckReadsOpencodeSkillsFromSharedDirectories(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer")
	putParityFile(fileSystem, ".claude/skills/review/SKILL.md", "review")
	putParityFile(fileSystem, ".agents/skills/review/SKILL.md", "review")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckExemptsCodexFromCommands(t *testing.T) {
	fileSystem := fstest.MapFS{}
	putSubagents(fileSystem, "drill-reviewer")
	putParityFile(fileSystem, ".claude/commands/review.md", "review")

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 1 || findings[0].Harness != opencodeHarness {
		t.Fatalf("expected only opencode to be reported, got %#v", findings)
	}
	if notes := UnsupportedNotes(); len(notes) == 0 || !strings.Contains(notes[0], "Codex") {
		t.Fatalf("expected the Codex exemption to be reported, got %v", notes)
	}
}

func TestFindingMessageNamesAndPluralizesTheDifference(t *testing.T) {
	singular := Finding{Capability: "subagent", Harness: codexHarness, Missing: []string{"planner"}}.Message()
	if !strings.Contains(singular, "subagent") ||
		!strings.Contains(singular, codexHarness) ||
		!strings.Contains(singular, "planner") {
		t.Fatalf("expected the capability, harness, and entry, got %q", singular)
	}

	plural := Finding{Capability: "command", Harness: opencodeHarness, Missing: []string{"plan", "review"}}.Message()
	if !strings.Contains(plural, "commands") || !strings.Contains(plural, "plan, review") {
		t.Fatalf("expected a plural capability and both entries, got %q", plural)
	}
}

func TestCheckReportsInjectedHarnessDirectoryReadFailure(t *testing.T) {
	injectedErr := errors.New("injected read failure")
	fileSystem := failingFS{FS: fstest.MapFS{}, path: ".claude/agents", err: injectedErr}

	_, err := CheckFS(fileSystem)
	if err == nil || !strings.Contains(err.Error(), "read .claude/agents") || !errors.Is(err, injectedErr) {
		t.Fatalf("expected a contextualized harness directory read error, got %v", err)
	}
}

type failingFS struct {
	fs.FS

	path string
	err  error
}

func (fileSystem failingFS) Open(name string) (fs.File, error) {
	if name == fileSystem.path {
		return nil, fileSystem.err
	}
	file, err := fileSystem.FS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}
	return file, nil
}

func putSubagents(fileSystem fstest.MapFS, names ...string) {
	for _, name := range names {
		putParityFile(fileSystem, ".claude/agents/"+name+".md", name)
		putParityFile(fileSystem, ".codex/agents/"+name+".toml", name)
		putParityFile(fileSystem, ".opencode/agents/"+name+".md", name)
	}
}

func putParityFile(fileSystem fstest.MapFS, path, contents string) {
	fileSystem[path] = &fstest.MapFile{Data: []byte(contents)}
}
