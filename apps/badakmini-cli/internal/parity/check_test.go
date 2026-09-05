package parity

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/fs"
	"slices"
	"strings"
	"testing"
	"testing/fstest"
)

const (
	testSkill = `---
name: review
description: Review work.
---

# Review

Canonical skill.
`
	testSkillRoute = "Read `.agents/skills/review/SKILL.md` completely, resolve every relative resource " +
		"from that skill directory, and follow it as authoritative before acting."
	testSkillAdapter = `---
name: review
description: Review work.
---

` + testSkillRoute + "\n"
	testAgent = `---
name: review
description: Review work.
mode: subagent
requires:
  - repository-read
denies:
  - repository-write
  - nested-agent
constraints:
  - inline-result-only
---

# Review

Canonical agent.
`
	testRoute = "Before acting, read the complete canonical agent definition at `.agents/agents/review.md` " +
		"from the repository root and follow it as authoritative. If it cannot be read, " +
		"stop and report the missing path."
)

func TestCheckAcceptsCanonicalContractAndReportsDigest(t *testing.T) {
	report, err := CheckFS(validContract())
	if err != nil {
		t.Fatalf("check valid contract: %v", err)
	}
	validCounts := report.Harnesses == harnessCount && report.Skills == 1 && report.Agents == 1
	if len(report.Findings) != 0 || !validCounts || len(report.Digest) != sha256.Size*2 {
		t.Fatalf("unexpected report: %#v", report)
	}
}

func TestCheckRejectsInstructionOverlays(t *testing.T) {
	fileSystem := validContract()
	fileSystem["CLAUDE.md"] = file("@AGENTS.md\nextra\n")
	fileSystem["nested/AGENTS.md"] = file("overlay")
	fileSystem["opencode.json"] = file(`{"instructions":["extra.md"]}`)
	report := checkReport(t, fileSystem)
	assertKinds(t, report, "invalid-instruction-adapter", "unexpected-instruction-source")
}

func TestCheckRejectsSkillAndAgentAdapterDrift(t *testing.T) {
	fileSystem := validContract()
	delete(fileSystem, ".claude/skills/review/SKILL.md")
	fileSystem[".opencode/agents/review.md"] = file(opencodeAgentAdapter("allow"))
	fileSystem[".claude/agents/extra.md"] = file("extra")
	report := checkReport(t, fileSystem)
	assertKinds(t, report, "agent-semantic-divergence", "missing-skill-adapter", "unexpected-agent-adapter")
}

func TestCanonicalSupportingResourceChangesDigest(t *testing.T) {
	fileSystem := validContract()
	before := checkReport(t, fileSystem).Digest
	fileSystem[".agents/skills/review/references/details.md"] = file("supporting resource")
	after := checkReport(t, fileSystem).Digest
	if before == after {
		t.Fatal("expected supporting resource to change digest")
	}
}

func TestFindingsAreStableSorted(t *testing.T) {
	fileSystem := validContract()
	delete(fileSystem, ".claude/agents/review.md")
	delete(fileSystem, ".claude/skills/review/SKILL.md")
	first := checkReport(t, fileSystem).Findings
	second := checkReport(t, fileSystem).Findings
	if !slices.EqualFunc(first, second, func(left, right Finding) bool { return left.Message() == right.Message() }) {
		t.Fatalf("expected stable findings, got %#v then %#v", first, second)
	}
	messages := make([]string, 0, len(first))
	for _, finding := range first {
		messages = append(messages, finding.Message())
	}
	if !slices.IsSorted(messages) {
		t.Fatalf("expected sorted findings, got %v", messages)
	}
}

func TestCheckReportsUnreadableCanonicalSource(t *testing.T) {
	injected := errors.New("injected failure")
	_, err := CheckFS(failingFS{FS: validContract(), path: "AGENTS.md", err: injected})
	if err == nil || !strings.Contains(err.Error(), "AGENTS.md") || !errors.Is(err, injected) {
		t.Fatalf("expected contextual error, got %v", err)
	}
}

func TestCheckReportsContextualStageErrors(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		prepare func(fstest.MapFS)
		want    string
	}{
		{name: "Claude instruction", path: "CLAUDE.md", want: "CLAUDE.md"},
		{name: "opencode config", path: "opencode.json", want: "opencode.json"},
		{name: "instruction walk", path: ".", want: "instruction sources"},
		{name: "skill directory", path: ".agents/skills", want: ".agents/skills"},
		{name: "skill manifest", path: ".agents/skills/review/SKILL.md", want: "SKILL.md"},
		{name: "skill adapter directory", path: ".claude/skills", want: ".claude/skills"},
		{name: "skill adapter", path: ".claude/skills/review/SKILL.md", want: "SKILL.md"},
		{name: "agent directory", path: ".agents/agents", want: ".agents/agents"},
		{name: "agent manifest", path: ".agents/agents/review.md", want: "review.md"},
		{name: "agent adapter", path: ".claude/agents/review.md", want: "review.md"},
		{
			name: "skill resource",
			path: ".agents/skills/review/references/details.md",
			prepare: func(fileSystem fstest.MapFS) {
				fileSystem[".agents/skills/review/references/details.md"] = file("details")
			},
			want: "canonical skill",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fileSystem := validContract()
			if test.prepare != nil {
				test.prepare(fileSystem)
			}
			_, err := CheckFS(failingFS{FS: fileSystem, path: test.path, err: errors.New("injected")})
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %q contextual error, got %v", test.want, err)
			}
		})
	}
}

func TestCheckHandlesOptionalAndMalformedEntries(t *testing.T) {
	fileSystem := validContract()
	fileSystem[".git/nested/AGENTS.md"] = file("ignored")
	fileSystem[".agents/skills/README.md"] = file("index")
	fileSystem[".agents/skills/empty"] = directory()
	fileSystem[".agents/skills/broken/SKILL.md"] = file("missing frontmatter")
	fileSystem[".agents/skills/mismatch/SKILL.md"] = file(`---
name: other
description: Other.
---
`)
	fileSystem[".claude/skills/README.md"] = file("index")
	fileSystem[".claude/skills/empty"] = directory()
	fileSystem[".claude/skills/extra/SKILL.md"] = file(testSkillAdapter)
	fileSystem[".claude/skills/review/SKILL.md"] = file("stale")
	fileSystem[".opencode/skills/extra/SKILL.md"] = file(testSkill)
	fileSystem[".agents/agents/README.md"] = file("index")
	fileSystem[".agents/agents/ignored.txt"] = file("ignored")
	fileSystem[".agents/agents/empty"] = directory()
	fileSystem[".agents/agents/broken.md"] = file("missing frontmatter")
	fileSystem[".claude/agents/README.md"] = file("index")
	fileSystem[".codex/agents/ignored.txt"] = file("ignored")
	fileSystem[".opencode/agents/empty"] = directory()

	report := checkReport(t, fileSystem)
	assertKinds(t, report, "invalid-agent", "invalid-skill", "skill-content-divergence", "unexpected-skill-adapter")
	for _, finding := range report.Findings {
		if finding.Path == ".git/nested/AGENTS.md" {
			t.Fatalf("ignored tree produced finding: %#v", finding)
		}
	}
}

//nolint:cyclop // This test exercises independent parser and adapter rejection branches.
func TestInstructionAndAdapterParsingBranches(t *testing.T) {
	t.Run("invalid opencode JSON", func(t *testing.T) {
		fileSystem := validContract()
		fileSystem["opencode.json"] = file("{")
		if _, err := CheckFS(fileSystem); err == nil || !strings.Contains(err.Error(), "invalid JSON") {
			t.Fatalf("expected invalid JSON error, got %v", err)
		}
	})

	t.Run("optional opencode config", func(t *testing.T) {
		fileSystem := validContract()
		delete(fileSystem, "opencode.json")
		if report := checkReport(t, fileSystem); len(report.Findings) != 0 {
			t.Fatalf("unexpected findings: %#v", report.Findings)
		}
	})

	if _, err := parseManifest("no frontmatter"); err == nil {
		t.Fatal("expected missing frontmatter error")
	}
	if _, err := parseManifest("---\nno terminator"); err == nil {
		t.Fatal("expected unterminated frontmatter error")
	}
	parsed, err := parseManifest("---\nignored\npermission:\n  edit: deny\nempty:\n  \n---\nbody")
	if err != nil || parsed.fields["permission.edit"] != permissionDeny {
		t.Fatalf("unexpected parsed manifest: %#v, %v", parsed, err)
	}
	if validSkillAdapter("invalid", canonicalEntry{}) {
		t.Fatal("invalid skill adapter passed")
	}
	extraSkillField := strings.Replace(
		testSkillAdapter,
		"description: Review work.",
		"description: Review work.\nprompt: extra",
		1,
	)
	if validSkillAdapter(extraSkillField, canonicalEntry{name: "review", description: "Review work."}) {
		t.Fatal("skill adapter with prompt-extending metadata passed")
	}
	if agentAdapterValid(claudeHarness, "invalid", canonicalEntry{}, "route") {
		t.Fatal("invalid agent adapter passed")
	}
	canonical := canonicalEntry{name: "review", description: "Review work."}
	badName := strings.Replace(claudeAgentAdapter(), "name: review", "name: other", 1)
	if agentAdapterValid(claudeHarness, badName, canonical, testRoute) {
		t.Fatal("Claude adapter with wrong name passed")
	}
	extraClaudeField := strings.Replace(claudeAgentAdapter(), "model: inherit", "model: inherit\nprompt: extra", 1)
	if agentAdapterValid(claudeHarness, extraClaudeField, canonical, testRoute) {
		t.Fatal("Claude adapter with prompt-extending metadata passed")
	}
	if agentAdapterValid(codexHarness, "name = invalid", canonical, testRoute) {
		t.Fatal("invalid Codex adapter passed")
	}
	if agentAdapterValid(codexHarness, codexAgentAdapter()+"prompt = \"extra\"\n", canonical, testRoute) {
		t.Fatal("Codex adapter with prompt-extending metadata passed")
	}
	multiline, err := parseManifest("---\ndescription:\n  continued value\n---\nbody")
	if err != nil || multiline.fields["description"] != "continued value" {
		t.Fatalf("expected multiline value, got %#v, %v", multiline, err)
	}
}

func TestDirectStageErrorBranches(t *testing.T) {
	injected := errors.New("injected")
	report := Report{}
	fileSystem := validContract()

	if _, _, err := readCanonicalSkills(
		failingFS{FS: fileSystem, path: ".agents/skills", err: injected},
		&report,
	); err == nil {
		t.Fatal("expected canonical skill directory error")
	}
	if err := inspectSkillAdapters(
		failingFS{FS: fileSystem, path: ".claude/skills", err: injected},
		&report,
		map[string]canonicalEntry{},
	); err == nil {
		t.Fatal("expected skill adapter directory error")
	}
	if err := inspectProhibitedSkillCopies(
		failingFS{FS: fileSystem, path: ".codex/skills", err: injected},
		&report,
	); err == nil {
		t.Fatal("expected prohibited skill directory error")
	}
	copyFS := fstest.MapFS{".opencode/skills/review/placeholder": file("placeholder")}
	if err := inspectProhibitedSkillCopies(
		failingFS{FS: copyFS, path: ".opencode/skills/review/SKILL.md", err: injected},
		&report,
	); err == nil {
		t.Fatal("expected prohibited skill manifest error")
	}
	if _, _, err := readCanonicalAgents(
		failingFS{FS: fileSystem, path: ".agents/agents", err: injected},
		&report,
	); err == nil {
		t.Fatal("expected canonical agent directory error")
	}

	emptyAgents := fstest.MapFS{
		".agents/agents":   directory(),
		".claude/agents":   directory(),
		".codex/agents":    directory(),
		".opencode/agents": directory(),
	}
	if _, _, err := inspectAgents(
		failingFS{FS: emptyAgents, path: ".claude/agents", err: injected},
		&report,
	); err == nil {
		t.Fatal("expected unexpected-adapter directory error")
	}

	resourceFS := validContract()
	resourceFS[".agents/skills/review/references/details.md"] = file("details")
	if err := addSkillDigest(
		failingFS{FS: resourceFS, path: ".agents/skills/review/references", err: injected},
		"review",
		map[string]string{},
	); err == nil {
		t.Fatal("expected skill walk entry error")
	}
}

//nolint:funlen // One table plus focused assertions covers the complete native capability mapping.
func TestNativePermissionMappings(t *testing.T) {
	tests := []struct {
		name    string
		denial  string
		harness string
		fields  map[string]string
		want    bool
	}{
		{name: "unmapped denial", denial: "other", harness: claudeHarness, want: true},
		{
			name: "Claude write denied", denial: "repository-write", harness: claudeHarness,
			fields: map[string]string{"tools": "Read, Edit"},
		},
		{
			name: "opencode write denied", denial: "repository-write", harness: opencodeHarness,
			fields: map[string]string{"permission.edit": "allow"},
		},
		{name: "Codex write denial stays canonical", denial: "repository-write", harness: codexHarness, want: true},
		{
			name: "Claude shell denied", denial: "shell", harness: claudeHarness,
			fields: map[string]string{"tools": "Read, Bash"},
		},
		{
			name: "opencode shell denied", denial: "shell", harness: opencodeHarness,
			fields: map[string]string{"permission.bash": "ask"},
		},
		{name: "Codex shell denial stays canonical", denial: "shell", harness: codexHarness, want: true},
		{
			name: "Claude nested agent denied", denial: "nested-agent", harness: claudeHarness,
			fields: map[string]string{"tools": "Read, Task"},
		},
		{
			name: "opencode nested agent denied", denial: "nested-agent", harness: opencodeHarness,
			fields: map[string]string{"permission.task": "allow"},
		},
		{name: "Codex nested denial stays canonical", denial: "nested-agent", harness: codexHarness, want: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			canonical := canonicalEntry{manifest: manifest{lists: map[string][]string{"denies": {test.denial}}}}
			adapter := manifest{fields: test.fields}
			if got := markdownPermissionsPreserve(canonical, adapter, test.harness); got != test.want {
				t.Fatalf("expected %t, got %t", test.want, got)
			}
		})
	}
	if requiredCapabilitiesPreserved([]string{"repository-read"}, manifest{fields: map[string]string{}}, claudeHarness) {
		t.Fatal("Claude adapter without required read capability passed")
	}
	if requiredCapabilitiesPreserved(
		[]string{"approval-or-read-only-shell"},
		manifest{fields: map[string]string{"permission.bash": "deny"}},
		opencodeHarness,
	) {
		t.Fatal("opencode adapter without required approval-gated shell passed")
	}
	if !requiredCapabilitiesPreserved(
		[]string{"approval-or-read-only-shell"},
		manifest{fields: map[string]string{"tools": "Read, Bash"}},
		claudeHarness,
	) {
		t.Fatal("Claude adapter with required shell capability failed")
	}
}

//nolint:cyclop // The compact helper matrix covers every small parsing branch.
func TestSmallParsingHelpers(t *testing.T) {
	finding := Finding{Kind: "kind", Path: "path", Field: "field", Problem: "problem"}
	if message := finding.Message(); !strings.Contains(message, "path:field") {
		t.Fatalf("field missing from %q", message)
	}
	for _, value := range []any{nil, "", []any{}} {
		if !emptyJSONValue(value) {
			t.Fatalf("expected empty value: %#v", value)
		}
	}
	for _, value := range []any{"present", []any{1}, map[string]any{}} {
		if emptyJSONValue(value) {
			t.Fatalf("expected non-empty value: %#v", value)
		}
	}
	if tomlScalar("name = invalid", "name") != "" || tomlScalar("other = \"x\"", "name") != "" {
		t.Fatal("invalid or missing TOML scalar passed")
	}
	if tomlMultiline("other = \"\"\"x\"\"\"", "name") != "" ||
		tomlMultiline("name = \"\"\"unterminated", "name") != "" {
		t.Fatal("invalid or missing TOML multiline passed")
	}
	if !containsWord("Read, Edit", "Edit") {
		t.Fatal("expected comma-separated word match")
	}
	validTOML := "name = \"review\"\n\ndeveloper_instructions = \"\"\"\nroute\n\"\"\"\n"
	if !tomlFieldsAllowed(validTOML, "name", "developer_instructions") {
		t.Fatal("expected allowed TOML fields")
	}
	for _, invalid := range []string{
		"name = \"review\"\nextra = \"value\"\n",
		"name = \"review\"\nname = \"again\"\n",
		"name = \"review\"\ndeveloper_instructions = \"\"\"\nunterminated\n",
		"not an assignment\n",
	} {
		if tomlFieldsAllowed(invalid, "name", "developer_instructions") {
			t.Fatalf("invalid TOML fields passed: %q", invalid)
		}
	}
}

func validContract() fstest.MapFS {
	return fstest.MapFS{
		"AGENTS.md":                      file("canonical rules\n"),
		"CLAUDE.md":                      file("@AGENTS.md\n"),
		"opencode.json":                  file(`{"$schema":"https://opencode.ai/config.json"}`),
		".agents/skills/review/SKILL.md": file(testSkill),
		".claude/skills/review/SKILL.md": file(testSkillAdapter),
		".agents/agents/review.md":       file(testAgent),
		".claude/agents/review.md":       file(claudeAgentAdapter()),
		".codex/agents/review.toml":      file(codexAgentAdapter()),
		".opencode/agents/review.md":     file(opencodeAgentAdapter(permissionDeny)),
	}
}

func claudeAgentAdapter() string {
	return `---
name: review
description: Review work.
tools: Read
model: inherit
---

` + testRoute + "\n"
}

func codexAgentAdapter() string {
	return `name = "review"
description = "Review work."
sandbox_mode = "read-only"

developer_instructions = """
` + testRoute + `
"""
`
}

func opencodeAgentAdapter(editPermission string) string {
	return `---
description: Review work.
mode: subagent
permission:
  edit: ` + editPermission + `
  task: deny
---

` + testRoute + "\n"
}

func checkReport(t *testing.T, fileSystem fs.FS) Report {
	t.Helper()
	report, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("check contract: %v", err)
	}
	return report
}

func assertKinds(t *testing.T, report Report, expected ...string) {
	t.Helper()
	kinds := make([]string, 0, len(report.Findings))
	for _, finding := range report.Findings {
		kinds = append(kinds, finding.Kind)
	}
	for _, wanted := range expected {
		if !slices.Contains(kinds, wanted) {
			t.Fatalf("expected %s in %v", wanted, kinds)
		}
	}
}

func file(contents string) *fstest.MapFile {
	return &fstest.MapFile{Data: []byte(contents)}
}

func directory() *fstest.MapFile {
	return &fstest.MapFile{Mode: fs.ModeDir}
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
	opened, err := fileSystem.FS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}
	return opened, nil
}
