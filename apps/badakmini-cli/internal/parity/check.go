// Package parity validates the canonical repository contract shared by every harness.
package parity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"path"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	claudeHarness   = "Claude Code"
	codexHarness    = "Codex"
	opencodeHarness = "opencode"
	harnessCount    = 3
	indexName       = "README.md"
	markdownExt     = ".md"
	permissionDeny  = "deny"
	pairSize        = 2
)

var skippedTrees = map[string]bool{
	".git": true, ".nx": true, "node_modules": true, "local-tmp": true, "generated-reports": true,
}

// Finding describes one deterministic contract violation.
type Finding struct {
	Kind       string
	Path       string
	Field      string
	Problem    string
	Capability string
	Harness    string
	Missing    []string
}

// Message formats both current contract findings and legacy injected test findings.
func (finding Finding) Message() string {
	if finding.Problem != "" {
		location := finding.Path
		if finding.Field != "" {
			location += ":" + finding.Field
		}
		return fmt.Sprintf("%s: %s: %s", finding.Kind, location, finding.Problem)
	}
	return fmt.Sprintf(
		"%s parity: %s is missing %s.",
		finding.Capability,
		finding.Harness,
		strings.Join(finding.Missing, ", "),
	)
}

// Report is the complete read-only contract result.
type Report struct {
	Findings  []Finding
	Harnesses int
	Skills    int
	Agents    int
	Digest    string
}

type manifest struct {
	fields map[string]string
	lists  map[string][]string
	body   string
}

type canonicalEntry struct {
	name        string
	description string
	manifest    manifest
}

// CheckFS inspects canonical sources and every required native adapter.
func CheckFS(fileSystem fs.FS) (Report, error) {
	report := Report{Harnesses: harnessCount}
	digestEntries := map[string]string{}

	agentsBody, err := readRequired(fileSystem, "AGENTS.md")
	if err != nil {
		return Report{}, err
	}
	digestEntries["AGENTS.md"] = normalizeText(agentsBody)

	if err := inspectInstructions(fileSystem, &report); err != nil {
		return Report{}, err
	}

	skills, skillDigest, err := inspectSkills(fileSystem, &report)
	if err != nil {
		return Report{}, err
	}
	report.Skills = len(skills)
	maps.Copy(digestEntries, skillDigest)

	agents, agentDigest, err := inspectAgents(fileSystem, &report)
	if err != nil {
		return Report{}, err
	}
	report.Agents = len(agents)
	maps.Copy(digestEntries, agentDigest)

	report.Digest = digest(digestEntries)
	sort.Slice(report.Findings, func(first, second int) bool {
		left, right := report.Findings[first], report.Findings[second]
		return left.Kind+"\x00"+left.Path+"\x00"+left.Field+"\x00"+left.Problem <
			right.Kind+"\x00"+right.Path+"\x00"+right.Field+"\x00"+right.Problem
	})
	return report, nil
}

func inspectInstructions(fileSystem fs.FS, report *Report) error {
	if err := inspectClaudeInstruction(fileSystem, report); err != nil {
		return err
	}
	if err := inspectOpenCodeInstructions(fileSystem, report); err != nil {
		return err
	}
	if err := inspectNestedInstructions(fileSystem, report); err != nil {
		return fmt.Errorf("inspect instruction sources: %w", err)
	}
	return nil
}

func inspectClaudeInstruction(fileSystem fs.FS, report *Report) error {
	claude, err := readRequired(fileSystem, "CLAUDE.md")
	if err != nil {
		return err
	}
	if strings.TrimSpace(strings.TrimPrefix(claude, "\ufeff")) != "@AGENTS.md" {
		addFinding(report, "invalid-instruction-adapter", "CLAUDE.md", "content", "must contain only @AGENTS.md")
	}
	return nil
}

func inspectOpenCodeInstructions(fileSystem fs.FS, report *Report) error {
	contents, err := fs.ReadFile(fileSystem, "opencode.json")
	if err == nil {
		var config map[string]any
		if parseErr := json.Unmarshal(contents, &config); parseErr != nil {
			return fmt.Errorf("read opencode.json: invalid JSON: %w", parseErr)
		}
		if instructions, present := config["instructions"]; present && !emptyJSONValue(instructions) {
			addFinding(report, "unexpected-instruction-source", "opencode.json", "instructions", "must be empty")
		}
		return nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("read opencode.json: %w", err)
	}
	return nil
}

func inspectNestedInstructions(fileSystem fs.FS, report *Report) error {
	err := fs.WalkDir(fileSystem, ".", func(entryPath string, entry fs.DirEntry, walkErr error) error {
		return inspectInstructionEntry(report, entryPath, entry, walkErr)
	})
	if err != nil {
		return fmt.Errorf("walk repository instructions: %w", err)
	}
	return nil
}

func inspectInstructionEntry(report *Report, entryPath string, entry fs.DirEntry, walkErr error) error {
	if walkErr != nil {
		return fmt.Errorf("visit %s: %w", entryPath, walkErr)
	}
	if entry.IsDir() {
		if entryPath != "." && skippedTrees[path.Base(entryPath)] {
			return fs.SkipDir
		}
		return nil
	}
	if slices.Contains([]string{"AGENTS.md", "CLAUDE.md"}, entryPath) {
		return nil
	}
	if unexpectedInstructionPath(entryPath) {
		addFinding(report, "unexpected-instruction-source", entryPath, "", "only root AGENTS.md is canonical")
	}
	return nil
}

func unexpectedInstructionPath(entryPath string) bool {
	instructionNames := []string{"AGENTS.md", "AGENTS.override.md", "CLAUDE.md"}
	return slices.Contains(instructionNames, path.Base(entryPath)) || strings.HasPrefix(entryPath, ".claude/rules/")
}

func inspectSkills(fileSystem fs.FS, report *Report) (map[string]canonicalEntry, map[string]string, error) {
	skills, digestEntries, err := readCanonicalSkills(fileSystem, report)
	if err != nil {
		return nil, nil, err
	}
	if err := inspectSkillAdapters(fileSystem, report, skills); err != nil {
		return nil, nil, err
	}
	if err := inspectProhibitedSkillCopies(fileSystem, report); err != nil {
		return nil, nil, err
	}
	return skills, digestEntries, nil
}

func readCanonicalSkills(
	fileSystem fs.FS,
	report *Report,
) (map[string]canonicalEntry, map[string]string, error) {
	entries, err := readDirOptional(fileSystem, ".agents/skills")
	if err != nil {
		return nil, nil, err
	}
	skills := map[string]canonicalEntry{}
	digestEntries := map[string]string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		canonical, present, readErr := readCanonicalSkill(fileSystem, report, name)
		if readErr != nil {
			return nil, nil, readErr
		}
		if !present {
			continue
		}
		skills[name] = canonical
		if err := addSkillDigest(fileSystem, name, digestEntries); err != nil {
			return nil, nil, err
		}
	}
	return skills, digestEntries, nil
}

func readCanonicalSkill(
	fileSystem fs.FS,
	report *Report,
	name string,
) (canonicalEntry, bool, error) {
	manifestPath := path.Join(".agents/skills", name, "SKILL.md")
	contents, err := fs.ReadFile(fileSystem, manifestPath)
	if errors.Is(err, fs.ErrNotExist) {
		return canonicalEntry{}, false, nil
	}
	if err != nil {
		return canonicalEntry{}, false, fmt.Errorf("read %s: %w", manifestPath, err)
	}
	parsed, err := parseManifest(string(contents))
	if err != nil {
		addFinding(report, "invalid-skill", manifestPath, "frontmatter", err.Error())
		return canonicalEntry{}, false, nil
	}
	if parsed.fields["name"] != name || parsed.fields["description"] == "" {
		addFinding(report, "invalid-skill", manifestPath, "name", "directory, name, and description must agree")
		return canonicalEntry{}, false, nil
	}
	canonical := canonicalEntry{name: name, description: parsed.fields["description"], manifest: parsed}
	return canonical, true, nil
}

func addSkillDigest(fileSystem fs.FS, name string, entries map[string]string) error {
	root := path.Join(".agents/skills", name)
	err := fs.WalkDir(fileSystem, root, func(itemPath string, item fs.DirEntry, itemErr error) error {
		if itemErr != nil {
			return itemErr
		}
		if item.IsDir() || item.Type()&fs.ModeSymlink != 0 {
			return nil
		}
		body, err := fs.ReadFile(fileSystem, itemPath)
		if err != nil {
			return fmt.Errorf("read %s: %w", itemPath, err)
		}
		entries[itemPath] = normalizeText(string(body))
		return nil
	})
	if err != nil {
		return fmt.Errorf("read canonical skill %s: %w", name, err)
	}
	return nil
}

func inspectSkillAdapters(fileSystem fs.FS, report *Report, skills map[string]canonicalEntry) error {
	adapterEntries, err := readDirOptional(fileSystem, ".claude/skills")
	if err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, entry := range adapterEntries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		adapterPath := path.Join(".claude/skills", name, "SKILL.md")
		contents, readErr := fs.ReadFile(fileSystem, adapterPath)
		if errors.Is(readErr, fs.ErrNotExist) {
			continue
		}
		if readErr != nil {
			return fmt.Errorf("read %s: %w", adapterPath, readErr)
		}
		seen[name] = true
		canonical, present := skills[name]
		if !present {
			addFinding(report, "unexpected-skill-adapter", adapterPath, "", "no canonical skill exists")
			continue
		}
		if !validSkillAdapter(string(contents), canonical) {
			addFinding(report, "skill-content-divergence", adapterPath, "", "description and canonical route must match")
		}
	}
	for name := range skills {
		if !seen[name] {
			adapterPath := path.Join(".claude/skills", name, "SKILL.md")
			addFinding(report, "missing-skill-adapter", adapterPath, "", "Claude adapter is required")
		}
	}
	return nil
}

func validSkillAdapter(contents string, canonical canonicalEntry) bool {
	parsed, err := parseManifest(contents)
	if err != nil {
		return false
	}
	expectedBody := fmt.Sprintf(
		"Read `.agents/skills/%s/SKILL.md` completely, resolve every relative resource "+
			"from that skill directory, and follow it as authoritative before acting.",
		canonical.name,
	)
	return manifestFieldsAllowed(parsed, "name", "description") &&
		parsed.fields["name"] == canonical.name &&
		parsed.fields["description"] == canonical.description &&
		collapseWhitespace(parsed.body) == expectedBody
}

func inspectProhibitedSkillCopies(fileSystem fs.FS, report *Report) error {
	for _, directory := range []string{".codex/skills", ".opencode/skills"} {
		entries, err := readDirOptional(fileSystem, directory)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			manifestPath := path.Join(directory, entry.Name(), "SKILL.md")
			_, err := fs.Stat(fileSystem, manifestPath)
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			if err != nil {
				return fmt.Errorf("inspect %s: %w", manifestPath, err)
			}
			addFinding(report, "unexpected-skill-adapter", manifestPath, "", "canonical bundle must be used directly")
		}
	}
	return nil
}

func inspectAgents(fileSystem fs.FS, report *Report) (map[string]canonicalEntry, map[string]string, error) {
	agents, digestEntries, err := readCanonicalAgents(fileSystem, report)
	if err != nil {
		return nil, nil, err
	}
	for _, canonical := range agents {
		if err := inspectAgentAdapters(fileSystem, report, canonical); err != nil {
			return nil, nil, err
		}
	}
	if err := inspectUnexpectedAgentAdapters(fileSystem, report, agents); err != nil {
		return nil, nil, err
	}
	return agents, digestEntries, nil
}

func readCanonicalAgents(
	fileSystem fs.FS,
	report *Report,
) (map[string]canonicalEntry, map[string]string, error) {
	entries, err := readDirOptional(fileSystem, ".agents/agents")
	if err != nil {
		return nil, nil, err
	}
	agents := map[string]canonicalEntry{}
	digestEntries := map[string]string{}
	for _, entry := range entries {
		if entry.IsDir() || path.Ext(entry.Name()) != markdownExt || entry.Name() == indexName {
			continue
		}
		entryPath := path.Join(".agents/agents", entry.Name())
		contents, readErr := fs.ReadFile(fileSystem, entryPath)
		if readErr != nil {
			return nil, nil, fmt.Errorf("read %s: %w", entryPath, readErr)
		}
		parsed, parseErr := parseManifest(string(contents))
		name := strings.TrimSuffix(entry.Name(), markdownExt)
		if parseErr != nil || !validCanonicalAgent(parsed, name) {
			problem := "name, description, mode, requires, denies, and constraints are required"
			addFinding(report, "invalid-agent", entryPath, "frontmatter", problem)
			continue
		}
		agents[name] = canonicalEntry{name: name, description: parsed.fields["description"], manifest: parsed}
		digestEntries[entryPath] = normalizeText(string(contents))
	}
	return agents, digestEntries, nil
}

func validCanonicalAgent(parsed manifest, name string) bool {
	return parsed.fields["name"] == name && parsed.fields["description"] != "" &&
		parsed.fields["mode"] == "subagent" && len(parsed.lists["requires"]) > 0 &&
		len(parsed.lists["denies"]) > 0 && len(parsed.lists["constraints"]) > 0
}

func inspectUnexpectedAgentAdapters(
	fileSystem fs.FS,
	report *Report,
	agents map[string]canonicalEntry,
) error {
	for _, adapter := range []struct{ directory, extension string }{
		{".claude/agents", markdownExt}, {".codex/agents", ".toml"}, {".opencode/agents", markdownExt},
	} {
		adapterEntries, readErr := readDirOptional(fileSystem, adapter.directory)
		if readErr != nil {
			return readErr
		}
		for _, entry := range adapterEntries {
			if entry.IsDir() || path.Ext(entry.Name()) != adapter.extension || entry.Name() == indexName {
				continue
			}
			name := strings.TrimSuffix(entry.Name(), adapter.extension)
			if _, present := agents[name]; !present {
				adapterPath := path.Join(adapter.directory, entry.Name())
				addFinding(report, "unexpected-agent-adapter", adapterPath, "", "no canonical agent exists")
			}
		}
	}
	return nil
}

func inspectAgentAdapters(fileSystem fs.FS, report *Report, canonical canonicalEntry) error {
	route := canonicalAgentRoute(canonical.name)
	adapters := []struct {
		harness, adapterPath string
	}{
		{claudeHarness, path.Join(".claude/agents", canonical.name+".md")},
		{codexHarness, path.Join(".codex/agents", canonical.name+".toml")},
		{opencodeHarness, path.Join(".opencode/agents", canonical.name+".md")},
	}
	for _, adapter := range adapters {
		contents, err := fs.ReadFile(fileSystem, adapter.adapterPath)
		if errors.Is(err, fs.ErrNotExist) {
			addFinding(report, "missing-agent-adapter", adapter.adapterPath, "", adapter.harness+" adapter is required")
			continue
		}
		if err != nil {
			return fmt.Errorf("read %s: %w", adapter.adapterPath, err)
		}
		if !agentAdapterValid(adapter.harness, string(contents), canonical, route) {
			problem := "description, route, or native permissions diverge"
			addFinding(report, "agent-semantic-divergence", adapter.adapterPath, "", problem)
		}
	}
	return nil
}

func canonicalAgentRoute(name string) string {
	return fmt.Sprintf(
		"Before acting, read the complete canonical agent definition at `.agents/agents/%s.md` "+
			"from the repository root and follow it as authoritative. If it cannot be read, stop and report the missing path.",
		name,
	)
}

func agentAdapterValid(harness, contents string, canonical canonicalEntry, route string) bool {
	if harness == codexHarness {
		return codexAgentAdapterValid(contents, canonical, route)
	}
	parsed, err := parseManifest(contents)
	if err != nil || parsed.fields["description"] != canonical.description ||
		collapseWhitespace(parsed.body) != route {
		return false
	}
	if harness == claudeHarness {
		return claudeAgentAdapterValid(parsed, canonical)
	}
	return opencodeAgentAdapterValid(parsed, canonical)
}

func codexAgentAdapterValid(contents string, canonical canonicalEntry, route string) bool {
	return tomlFieldsAllowed(contents, "name", "description", "sandbox_mode", "developer_instructions") &&
		tomlScalar(contents, "name") == canonical.name &&
		tomlScalar(contents, "description") == canonical.description &&
		tomlScalar(contents, "sandbox_mode") == "read-only" &&
		strings.TrimSpace(tomlMultiline(contents, "developer_instructions")) == route
}

func claudeAgentAdapterValid(parsed manifest, canonical canonicalEntry) bool {
	return manifestFieldsAllowed(parsed, "name", "description", "tools", "model") &&
		parsed.fields["name"] == canonical.name && parsed.fields["model"] == "inherit" &&
		markdownPermissionsPreserve(canonical, parsed, claudeHarness)
}

func opencodeAgentAdapterValid(parsed manifest, canonical canonicalEntry) bool {
	return manifestFieldsAllowed(
		parsed,
		"description",
		"mode",
		"permission",
		"permission.edit",
		"permission.bash",
		"permission.task",
	) && parsed.fields["mode"] == "subagent" && markdownPermissionsPreserve(canonical, parsed, opencodeHarness)
}

func manifestFieldsAllowed(parsed manifest, allowed ...string) bool {
	if len(parsed.lists) != 0 {
		return false
	}
	for field := range parsed.fields {
		if !slices.Contains(allowed, field) {
			return false
		}
	}
	return true
}

func markdownPermissionsPreserve(canonical canonicalEntry, adapter manifest, harness string) bool {
	denies := canonical.manifest.lists["denies"]
	return requiredCapabilitiesPreserved(canonical.manifest.lists["requires"], adapter, harness) &&
		repositoryWritePreserved(denies, adapter, harness) &&
		shellPreserved(denies, adapter, harness) && nestedAgentPreserved(denies, adapter, harness)
}

func requiredCapabilitiesPreserved(requires []string, adapter manifest, harness string) bool {
	if harness == claudeHarness && slices.Contains(requires, "repository-read") &&
		!containsWord(adapter.fields["tools"], "Read") {
		return false
	}
	if !slices.Contains(requires, "approval-or-read-only-shell") {
		return true
	}
	if harness == claudeHarness {
		return containsWord(adapter.fields["tools"], "Bash")
	}
	return harness != opencodeHarness || adapter.fields["permission.bash"] == "ask"
}

func repositoryWritePreserved(denies []string, adapter manifest, harness string) bool {
	if !slices.Contains(denies, "repository-write") {
		return true
	}
	if harness == claudeHarness {
		return !containsWord(adapter.fields["tools"], "Edit") && !containsWord(adapter.fields["tools"], "Write")
	}
	return harness != opencodeHarness || adapter.fields["permission.edit"] == permissionDeny
}

func shellPreserved(denies []string, adapter manifest, harness string) bool {
	if !slices.Contains(denies, "shell") {
		return true
	}
	if harness == claudeHarness {
		return !containsWord(adapter.fields["tools"], "Bash")
	}
	return harness != opencodeHarness || adapter.fields["permission.bash"] == permissionDeny
}

func nestedAgentPreserved(denies []string, adapter manifest, harness string) bool {
	if !slices.Contains(denies, "nested-agent") {
		return true
	}
	if harness == claudeHarness {
		return !containsWord(adapter.fields["tools"], "Task")
	}
	return harness != opencodeHarness || adapter.fields["permission.task"] == permissionDeny
}

func parseManifest(contents string) (manifest, error) {
	normalized := normalizeText(contents)
	frontmatter, present := strings.CutPrefix(normalized, "---\n")
	if !present {
		return manifest{}, errors.New("missing frontmatter")
	}
	parts := strings.SplitN(frontmatter, "\n---\n", pairSize)
	if len(parts) != pairSize {
		return manifest{}, errors.New("unterminated frontmatter")
	}
	result := manifest{fields: map[string]string{}, lists: map[string][]string{}, body: parts[1]}
	current := ""
	for line := range strings.SplitSeq(parts[0], "\n") {
		parseManifestLine(&result, &current, line)
	}
	return result, nil
}

func parseManifestLine(result *manifest, current *string, line string) {
	if item, present := strings.CutPrefix(line, "  - "); present && *current != "" {
		result.lists[*current] = append(result.lists[*current], strings.TrimSpace(item))
		return
	}
	if value, present := strings.CutPrefix(line, "  "); present && *current != "" {
		parseIndentedManifestValue(result, *current, strings.TrimSpace(value))
		return
	}
	key, value, present := strings.Cut(line, ":")
	if !present {
		return
	}
	*current = strings.TrimSpace(key)
	result.fields[*current] = strings.Trim(strings.TrimSpace(value), "\"")
}

func parseIndentedManifestValue(result *manifest, current, value string) {
	key, nestedValue, nested := strings.Cut(value, ":")
	if nested && current == "permission" {
		result.fields[current+"."+strings.TrimSpace(key)] = strings.TrimSpace(nestedValue)
		return
	}
	if value != "" {
		result.fields[current] = strings.TrimSpace(result.fields[current] + " " + value)
	}
}

func readRequired(fileSystem fs.FS, entryPath string) (string, error) {
	contents, err := fs.ReadFile(fileSystem, entryPath)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", entryPath, err)
	}
	return string(contents), nil
}

func readDirOptional(fileSystem fs.FS, directory string) ([]fs.DirEntry, error) {
	entries, err := fs.ReadDir(fileSystem, directory)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", directory, err)
	}
	return entries, nil
}

func addFinding(report *Report, kind, entryPath, field, problem string) {
	report.Findings = append(report.Findings, Finding{Kind: kind, Path: entryPath, Field: field, Problem: problem})
}

func digest(entries map[string]string) string {
	paths := make([]string, 0, len(entries))
	for entryPath := range entries {
		paths = append(paths, entryPath)
	}
	sort.Strings(paths)
	hash := sha256.New()
	for _, entryPath := range paths {
		_, _ = fmt.Fprintf(hash, "%s\x00%s\x00", entryPath, entries[entryPath])
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func normalizeText(value string) string {
	return strings.ReplaceAll(strings.TrimPrefix(value, "\ufeff"), "\r\n", "\n")
}

func emptyJSONValue(value any) bool {
	switch typed := value.(type) {
	case nil:
		return true
	case string:
		return typed == ""
	case []any:
		return len(typed) == 0
	default:
		return false
	}
}

func tomlFieldsAllowed(contents string, allowed ...string) bool {
	fields := []string{}
	inMultiline := false
	for line := range strings.SplitSeq(contents, "\n") {
		trimmed := strings.TrimSpace(line)
		if inMultiline {
			if trimmed == `"""` {
				inMultiline = false
			}
			continue
		}
		key, value, present := strings.Cut(trimmed, "=")
		if !present {
			continue
		}
		key = strings.TrimSpace(key)
		if !slices.Contains(allowed, key) || slices.Contains(fields, key) {
			return false
		}
		fields = append(fields, key)
		inMultiline = strings.TrimSpace(value) == `"""`
	}
	return !inMultiline && len(fields) == len(allowed)
}

func tomlScalar(contents, key string) string {
	prefix := key + " = "
	for line := range strings.SplitSeq(contents, "\n") {
		value, present := strings.CutPrefix(strings.TrimSpace(line), prefix)
		if !present {
			continue
		}
		unquoted, err := strconv.Unquote(strings.TrimSpace(value))
		if err == nil {
			return unquoted
		}
	}
	return ""
}

func tomlMultiline(contents, key string) string {
	marker := key + " = \"\"\""
	_, rest, present := strings.Cut(contents, marker)
	if !present {
		return ""
	}
	value, _, present := strings.Cut(rest, "\"\"\"")
	if !present {
		return ""
	}
	return value
}

func containsWord(value, wanted string) bool {
	for item := range strings.SplitSeq(value, ",") {
		if strings.TrimSpace(item) == wanted {
			return true
		}
	}
	return false
}

func collapseWhitespace(value string) string {
	return strings.Join(strings.Fields(value), " ")
}
