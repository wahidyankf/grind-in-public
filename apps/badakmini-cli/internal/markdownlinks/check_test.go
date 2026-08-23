package markdownlinks

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCheckAcceptsValidLinksThroughRuntimeDoubles(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"README.md": `
[Guide](docs/guide.md)
[Guide section](docs/guide.md#getting-started)
[Guide with title](docs/guide.md "Optional title")
[Setext section](docs/guide.md#setext-heading)
[Local section](#introduction)
[Website](https://example.com/docs)
[Email](mailto:hello@example.com)

# Introduction

~~~markdown
[Ignored](also-missing.md)
~~~
`,
		"docs/guide.md": "# Getting Started\n\nSetext Heading\n--------------\n",
	}

	findings, err := Check(root, memoryRuntime(root, files, "README.md", "docs/guide.md"))
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsMissingTargetsAnchorsAndEscapesThroughRuntimeDoubles(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"README.md": `
[Missing](docs/missing.md)
[Missing anchor](docs/guide.md#not-there)
[Outside](../outside.md)
[Symlink outside](docs/external.md)
`,
		"docs/guide.md":    "# Present\n",
		"docs/external.md": "# External\n",
	}
	runtime := memoryRuntime(root, files, "README.md", "docs/guide.md")
	runtime.EvalSymlinks = func(path string) (string, error) {
		if path == filepath.Join(root, "docs", "external.md") {
			return filepath.Join("root", "outside", "external.md"), nil
		}
		return filepath.Clean(path), nil
	}

	findings, err := Check(root, runtime)
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 4 {
		t.Fatalf("expected four findings, got %#v", findings)
	}
	assertFinding(t, findings[0], "docs/missing.md", "targets a file that does not exist")
	assertFinding(t, findings[1], "docs/guide.md#not-there", "targets a heading that does not exist")
	assertFinding(t, findings[2], "../outside.md", "points outside this repository")
	assertFinding(t, findings[3], "docs/external.md", "resolves outside this repository")
}

func TestCheckSupportsReferencesEncodingDirectoriesAndDuplicateAnchors(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"README.md": `
[Guide][guide]
[Directory](docs/#second-heading-1)

[guide]: docs/guide%20file.md
`,
		"docs/README.md":     "# First Heading\n# Second Heading\n# Second Heading\n",
		"docs/guide file.md": "# Guide\n",
	}

	findings, err := Check(root, memoryRuntime(root, files, "README.md"))
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckUsesOnlyTheInjectedTrackedSet(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"README.md":                       "# Repository\n",
		"node_modules/package/ignored.md": "[Broken](missing.md)\n",
		"apps/example/dist/ignored.md":    "[Broken](missing.md)\n",
	}

	findings, err := Check(root, memoryRuntime(root, files, "README.md", "notes.txt"))
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected untracked and non-Markdown files to be ignored, got %#v", findings)
	}
}

func TestCheckSupportsParserEdgeCasesThroughRuntimeDoubles(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"README.md": strings.Join([]string{
			"[Angle](<docs/guide file.md>)",
			"[Nested [label]](docs/nested_(guide).md)",
			"[Collapsed][]",
			"[Shortcut]",
			"[Titled](docs/titled.md\t\"Optional title\")",
			"`[Inline code](missing.md)`",
			`\[Escaped](missing.md)`,
			"",
			"[Collapsed]: docs/collapsed.md",
			"[Shortcut]: docs/shortcut.md",
		}, "\n"),
	}
	for _, path := range []string{
		"docs/guide file.md",
		"docs/nested_(guide).md",
		"docs/collapsed.md",
		"docs/shortcut.md",
		"docs/titled.md",
	} {
		files[path] = "# Present\n"
	}

	findings, err := Check(root, memoryRuntime(root, files, "README.md"))
	if err != nil {
		t.Fatalf("expected successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected parser edge cases to resolve, got %#v", findings)
	}
}

func TestCheckReportsInjectedDiscoveryAndReadFailures(t *testing.T) {
	root := filepath.Join("root", "repository")
	injectedErr := errors.New("injected failure")

	discoveryRuntime := memoryRuntime(root, map[string]string{})
	discoveryRuntime.TrackedFiles = func(string) (map[string]struct{}, error) {
		return nil, injectedErr
	}
	if _, err := Check(root, discoveryRuntime); !errors.Is(err, injectedErr) {
		t.Fatalf("expected tracked-file discovery failure, got %v", err)
	}

	readRuntime := memoryRuntime(root, map[string]string{}, "README.md")
	if _, err := Check(root, readRuntime); err == nil || !strings.Contains(err.Error(), "read README.md") {
		t.Fatalf("expected contextualized source read failure, got %v", err)
	}
}

func TestParseTrackedFilesPreservesNULDelimitedPaths(t *testing.T) {
	files := ParseTrackedFiles([]byte("README.md\x00docs/space name.md\x00\x00"))
	for _, path := range []string{"README.md", "docs/space name.md"} {
		if _, exists := files[path]; !exists {
			t.Fatalf("expected tracked path %q in %#v", path, files)
		}
	}
}

func TestCheckSortsFindingsByPathLineAndDestination(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"a.md": "[Z](missing-z.md) [A](missing-a.md)\n",
		"b.md": "[B](missing-b.md)\n",
	}

	findings, err := Check(root, memoryRuntime(root, files, "b.md", "a.md"))
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 3 || findings[0].Path != "a.md" || findings[0].Destination != "missing-a.md" ||
		findings[1].Destination != "missing-z.md" || findings[2].Path != "b.md" {
		t.Fatalf("expected stable finding order, got %#v", findings)
	}
}

func TestMalformedMarkdownLinksAreNotParsed(t *testing.T) {
	references := map[string]string{"known": "known.md"}
	tests := []struct {
		name     string
		line     string
		position int
	}{
		{name: "unclosed label", line: "[label", position: 0},
		{name: "unclosed inline destination", line: "[label](target", position: 0},
		{name: "unclosed full reference", line: "[label][known", position: 0},
		{name: "unknown shortcut reference", line: "[unknown]", position: 0},
	}

	for _, test := range tests {
		_, _, ok := parseLinkAt(test.line, test.position, 1, references)
		if ok {
			t.Fatalf("expected %s to be rejected", test.name)
		}
	}
}

func TestInlineDestinationBoundaries(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		opening     int
		destination string
		ok          bool
	}{
		{name: "only whitespace", value: "( \t", opening: 0, ok: false},
		{
			name: "angle destination", value: "(<path with spaces.md>)", opening: 0,
			destination: "path with spaces.md", ok: true,
		},
		{name: "angle missing closer", value: "(<path.md)", opening: 0, ok: false},
		{name: "angle missing parenthesis", value: "(<path.md>", opening: 0, ok: false},
		{name: "nested parentheses", value: "(path_(nested).md)", opening: 0, destination: "path_(nested).md", ok: true},
		{name: "escaped parenthesis", value: `(path\).md)`, opening: 0, destination: `path\).md`, ok: true},
		{name: "title without closer", value: `(path.md "title"`, opening: 0, ok: false},
	}

	for _, test := range tests {
		destination, _, parsed := inlineDestination(test.value, test.opening)
		if parsed != test.ok || destination != test.destination {
			t.Fatalf(
				"%s: got destination %q and ok=%t, want %q and ok=%t",
				test.name,
				destination,
				parsed,
				test.destination,
				test.ok,
			)
		}
	}
}

func TestMatchingBracketHandlesEscapesAndMissingCloser(t *testing.T) {
	value := `[outer \[literal\]]`
	if end, ok := matchingBracket(value, 0); !ok || end != len(value)-1 {
		t.Fatalf("expected escaped brackets to stay inside the label, got end=%d ok=%t", end, ok)
	}
	if _, ok := matchingBracket("[missing", 0); ok {
		t.Fatal("expected a missing closing bracket to fail")
	}
}

func TestCheckReportsTargetAndDestinationProblemsThroughDoubles(t *testing.T) {
	root := filepath.Join("root", "repository")
	files := map[string]string{
		"README.md": strings.Join([]string{
			"[Non-Markdown fragment](notes.txt#section)",
			"[Directory without README](empty/#section)",
			"[Invalid path](bad%zz.md)",
		}, "\n"),
		"notes.txt":   "plain text",
		"empty/.keep": "",
	}

	findings, err := Check(root, memoryRuntime(root, files, "README.md"))
	if err != nil {
		t.Fatalf("expected completed check, got %v", err)
	}
	if len(findings) != 3 {
		t.Fatalf("expected three target findings, got %#v", findings)
	}
	assertFinding(t, findings[0], "notes.txt#section", "uses a fragment on a non-Markdown target")
	assertFinding(t, findings[1], "empty/#section", "targets a directory without README.md for its fragment")
	assertFinding(t, findings[2], "bad%zz.md", "has an invalid URL")
}

func TestLocalTargetAndInspectionFailureBranches(t *testing.T) {
	root := filepath.Join("root", "repository")
	if got := localTargetPath(root, "docs/source.md", "/README.md"); got != filepath.Join(root, "README.md") {
		t.Fatalf("expected repository-root target, got %q", got)
	}

	injectedErr := errors.New("injected failure")
	tests := []struct {
		name    string
		runtime Runtime
		problem string
	}{
		{
			name: "stat failure",
			runtime: Runtime{Stat: func(string) (fs.FileInfo, error) {
				return nil, injectedErr
			}},
			problem: "cannot inspect its target",
		},
		{
			name: "root resolution failure",
			runtime: Runtime{
				Stat: func(string) (fs.FileInfo, error) { return memoryFileInfo{name: "target.md"}, nil },
				EvalSymlinks: func(string) (string, error) {
					return "", injectedErr
				},
			},
			problem: "cannot resolve the repository root",
		},
		{
			name: "target resolution failure",
			runtime: Runtime{
				Stat: func(string) (fs.FileInfo, error) { return memoryFileInfo{name: "target.md"}, nil },
				EvalSymlinks: func(path string) (string, error) {
					if path == root {
						return root, nil
					}
					return "", injectedErr
				},
			},
			problem: "cannot resolve its target",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, _, problem := inspectLocalTarget(test.runtime, root, filepath.Join(root, "target.md"))
			if problem != test.problem {
				t.Fatalf("got problem %q, want %q", problem, test.problem)
			}
		})
	}
}

func TestFragmentTargetFailureBranches(t *testing.T) {
	root := filepath.Join("root", "repository")
	fileInfo := memoryFileInfo{name: "target.md"}
	readFailureRuntime := Runtime{ReadFile: func(string) ([]byte, error) { return nil, errors.New("read failed") }}
	problem := checkFragmentTarget(
		readFailureRuntime,
		filepath.Join(root, "target.md"),
		"section",
		fileInfo,
		root,
	)
	if problem != "cannot read its fragment target" {
		t.Fatalf("expected fragment read failure, got %q", problem)
	}

	directoryInfo := memoryFileInfo{name: "docs", directory: true}
	tests := []struct {
		name    string
		runtime Runtime
		problem string
	}{
		{
			name: "README resolution failure",
			runtime: Runtime{
				Stat: func(string) (fs.FileInfo, error) { return fileInfo, nil },
				EvalSymlinks: func(string) (string, error) {
					return "", errors.New("resolve failed")
				},
			},
			problem: "cannot resolve its target",
		},
		{
			name: "README outside repository",
			runtime: Runtime{
				Stat: func(string) (fs.FileInfo, error) { return fileInfo, nil },
				EvalSymlinks: func(string) (string, error) {
					return filepath.Join("root", "outside", "README.md"), nil
				},
			},
			problem: "resolves outside this repository",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			problem := checkFragmentTarget(test.runtime, filepath.Join(root, "docs"), "section", directoryInfo, root)
			if problem != test.problem {
				t.Fatalf("got problem %q, want %q", problem, test.problem)
			}
		})
	}
}

func TestAnchorHelpersRejectEmptyAndMixedHeadings(t *testing.T) {
	if hasAnchor("# !!!\n", "anything") {
		t.Fatal("expected punctuation-only heading to produce no anchor")
	}
	if isSetextUnderline("-=-") {
		t.Fatal("expected mixed Setext markers to be rejected")
	}
}

func memoryRuntime(root string, files map[string]string, tracked ...string) Runtime {
	return Runtime{
		ReadFile: func(path string) ([]byte, error) {
			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return nil, fmt.Errorf("resolve read path: %w", err)
			}
			contents, ok := files[filepath.ToSlash(relativePath)]
			if !ok {
				return nil, &fs.PathError{Op: "read", Path: path, Err: fs.ErrNotExist}
			}
			return []byte(contents), nil
		},
		Stat: func(path string) (fs.FileInfo, error) {
			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return nil, fmt.Errorf("resolve stat path: %w", err)
			}
			relativePath = filepath.ToSlash(relativePath)
			if contents, ok := files[relativePath]; ok {
				return memoryFileInfo{name: filepath.Base(path), size: int64(len(contents))}, nil
			}
			prefix := strings.TrimSuffix(relativePath, "/") + "/"
			for filePath := range files {
				if strings.HasPrefix(filePath, prefix) {
					return memoryFileInfo{name: filepath.Base(path), directory: true}, nil
				}
			}
			return nil, &fs.PathError{Op: "stat", Path: path, Err: fs.ErrNotExist}
		},
		EvalSymlinks: func(path string) (string, error) {
			return filepath.Clean(path), nil
		},
		TrackedFiles: func(string) (map[string]struct{}, error) {
			result := make(map[string]struct{}, len(tracked))
			for _, path := range tracked {
				result[path] = struct{}{}
			}
			return result, nil
		},
	}
}

type memoryFileInfo struct {
	name      string
	size      int64
	directory bool
}

func (info memoryFileInfo) Name() string       { return info.name }
func (info memoryFileInfo) Size() int64        { return info.size }
func (info memoryFileInfo) ModTime() time.Time { return time.Time{} }
func (info memoryFileInfo) IsDir() bool        { return info.directory }
func (info memoryFileInfo) Sys() any           { return nil }
func (info memoryFileInfo) Mode() fs.FileMode {
	if info.directory {
		return fs.ModeDir | 0o750
	}
	return 0o600
}

func assertFinding(t *testing.T, finding Finding, destination, problem string) {
	t.Helper()
	if finding.Destination != destination || finding.Problem != problem {
		t.Fatalf("unexpected finding: %#v", finding)
	}
}
