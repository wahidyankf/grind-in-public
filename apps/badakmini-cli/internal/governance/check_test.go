package governance

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
)

func TestCheckAcceptsDocumentsAtTheWordLimit(t *testing.T) {
	fileSystem := newRepositoryFixture()
	putFile(fileSystem, agentsFile, words(MaxWords))
	putFile(fileSystem, claudeFile, words(MaxWords))
	putFile(fileSystem, governanceDirectory+"/policy.md", words(MaxWords))

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckReportsDocumentsOverTheWordLimit(t *testing.T) {
	fileSystem := newRepositoryFixture()
	putFile(fileSystem, agentsFile, words(MaxWords+1))
	putFile(fileSystem, claudeFile, words(MaxWords+1))
	putFile(fileSystem, governanceDirectory+"/nested/policy.md", words(MaxWords+1))
	putFile(fileSystem, governanceDirectory+"/nested/ignored.txt", words(MaxWords+1))

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 3 {
		t.Fatalf("expected three findings, got %#v", findings)
	}
	if findings[0].Path != agentsFile || findings[0].WordCount != MaxWords+1 {
		t.Fatalf("unexpected root finding: %#v", findings[0])
	}
	if findings[1].Path != claudeFile || findings[1].WordCount != MaxWords+1 {
		t.Fatalf("unexpected instruction file finding: %#v", findings[1])
	}
	if findings[2].Path != "repo-governance/nested/policy.md" {
		t.Fatalf("expected recursive Markdown finding, got %#v", findings[2])
	}
}

func TestCheckReportsOnlyAuthoredHarnessReadmes(t *testing.T) {
	fileSystem := newRepositoryFixture()
	putFile(fileSystem, ".claude/README.md", words(MaxWords+1))
	putFile(fileSystem, ".opencode/agents/README.md", words(MaxWords+1))
	putFile(fileSystem, ".opencode/agents/drill-reviewer.md", words(MaxWords+1))
	putFile(fileSystem, ".opencode/node_modules/dep/README.md", words(MaxWords+1))
	putFile(fileSystem, ".claude/.cache/README.md", words(MaxWords+1))

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 2 || findings[0].Path != ".claude/README.md" || findings[1].Path != ".opencode/agents/README.md" {
		t.Fatalf("expected only authored harness indexes, got %#v", findings)
	}
}

func TestCheckReportsTheSharedHarnessDirectory(t *testing.T) {
	fileSystem := newRepositoryFixture()
	putFile(fileSystem, ".agents/README.md", words(MaxWords+1))
	putFile(fileSystem, ".agents/skills/grill-me/SKILL.md", words(MaxWords+1))

	findings, err := CheckFS(fileSystem)
	if err != nil {
		t.Fatalf("expected a completed check, got %v", err)
	}
	if len(findings) != 1 || findings[0].Path != ".agents/README.md" {
		t.Fatalf("expected the shared harness index to be reported, got %#v", findings)
	}
}

func TestCheckAcceptsRepositoriesWithoutHarnessDirectories(t *testing.T) {
	findings, err := CheckFS(newRepositoryFixture())
	if err != nil {
		t.Fatalf("expected a successful check, got %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckRequiresGovernanceStructure(t *testing.T) {
	tests := []struct {
		name       string
		fileSystem fstest.MapFS
		message    string
	}{
		{name: "agents file", fileSystem: repositoryWith(claudeFile), message: "required file not found: AGENTS.md"},
		{name: "claude file", fileSystem: repositoryWith(agentsFile), message: "required file not found: CLAUDE.md"},
		{
			name: "governance directory",
			fileSystem: fstest.MapFS{
				agentsFile: &fstest.MapFile{Data: []byte("short")},
				claudeFile: &fstest.MapFile{Data: []byte("short")},
			},
			message: "required directory not found: repo-governance",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := CheckFS(test.fileSystem)
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Fatalf("expected %q, got %v", test.message, err)
			}
		})
	}
}

func TestCheckRejectsWrongGovernancePathTypes(t *testing.T) {
	tests := []struct {
		name       string
		fileSystem fstest.MapFS
		message    string
	}{
		{
			name: "directory in place of instruction file",
			fileSystem: fstest.MapFS{
				agentsFile:          &fstest.MapFile{Mode: fs.ModeDir},
				claudeFile:          &fstest.MapFile{Data: []byte("short")},
				governanceDirectory: &fstest.MapFile{Mode: fs.ModeDir},
			},
			message: "required file is not regular: AGENTS.md",
		},
		{
			name: "file in place of governance directory",
			fileSystem: fstest.MapFS{
				agentsFile:          &fstest.MapFile{Data: []byte("short")},
				claudeFile:          &fstest.MapFile{Data: []byte("short")},
				governanceDirectory: &fstest.MapFile{Data: []byte("short")},
			},
			message: "required path is not a directory: repo-governance",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := CheckFS(test.fileSystem)
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Fatalf("expected %q, got %v", test.message, err)
			}
		})
	}
}

//nolint:funlen // One table proves every filesystem failure keeps its policy context.
func TestFilesystemFailuresRetainPolicyContext(t *testing.T) {
	injectedErr := errors.New("injected filesystem failure")
	tests := []struct {
		name       string
		fileSystem fs.FS
		check      func(fs.FS) error
		message    string
	}{
		{
			name:       "document read",
			fileSystem: failingFS{FS: fstest.MapFS{}, path: "missing.md", err: injectedErr},
			check: func(fileSystem fs.FS) error {
				_, err := checkFile(fileSystem, "missing.md")
				return err
			},
			message: "read missing.md",
		},
		{
			name:       "harness inspection",
			fileSystem: failingFS{FS: fstest.MapFS{}, path: ".claude", err: injectedErr},
			check: func(fileSystem fs.FS) error {
				_, err := checkHarnessDirectory(fileSystem, ".claude")
				return err
			},
			message: "inspect .claude",
		},
		{
			name:       "instruction read",
			fileSystem: failingFS{FS: repositoryWith(agentsFile, claudeFile), path: agentsFile, err: injectedErr},
			check: func(fileSystem fs.FS) error {
				_, err := checkInstructionFiles(fileSystem)
				return err
			},
			message: "read AGENTS.md",
		},
		{
			name: "governance document read",
			fileSystem: failingFS{
				FS:   repositoryWith(agentsFile, claudeFile, "repo-governance/broken.md"),
				path: "repo-governance/broken.md",
				err:  injectedErr,
			},
			check: func(fileSystem fs.FS) error {
				_, err := CheckFS(fileSystem)
				return err
			},
			message: "read repo-governance/broken.md",
		},
		{
			name: "harness README read",
			fileSystem: failingFS{
				FS:   repositoryWith(agentsFile, claudeFile, ".agents/README.md"),
				path: ".agents/README.md",
				err:  injectedErr,
			},
			check: func(fileSystem fs.FS) error {
				_, err := CheckFS(fileSystem)
				return err
			},
			message: "read .agents/README.md",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.check(test.fileSystem)
			if err == nil || !strings.Contains(err.Error(), test.message) || !errors.Is(err, injectedErr) {
				t.Fatalf("expected contextualized injected failure, got %v", err)
			}
		})
	}
}

func TestRequiredPathsReportInspectionFailures(t *testing.T) {
	injectedErr := errors.New("inspect failed")
	fileSystem := failingFS{FS: fstest.MapFS{}, path: "broken", err: injectedErr}

	if err := requireFile(fileSystem, "broken"); err == nil || !strings.Contains(err.Error(), "inspect broken") {
		t.Fatalf("expected a file inspection error, got %v", err)
	}
	if err := requireDirectory(fileSystem, "broken"); err == nil || !strings.Contains(err.Error(), "inspect broken") {
		t.Fatalf("expected a directory inspection error, got %v", err)
	}
}

func TestVisitorsReturnFilesystemWalkErrors(t *testing.T) {
	walkErr := errors.New("walk failed")

	if err := visitGovernanceDocument(nil, "path", nil, walkErr, nil); !errors.Is(err, walkErr) {
		t.Fatalf("expected governance visitor to return the walk error, got %v", err)
	}
	if err := visitHarnessReadme(nil, "harness", "path", nil, walkErr, nil); !errors.Is(err, walkErr) {
		t.Fatalf("expected harness visitor to return the walk error, got %v", err)
	}
}

func TestVendoredDirectoryClassification(t *testing.T) {
	tests := []struct {
		name     string
		vendored bool
	}{
		{name: "node_modules", vendored: true},
		{name: ".cache", vendored: true},
		{name: "agents", vendored: false},
	}

	for _, test := range tests {
		if got := isVendoredDirectory(test.name); got != test.vendored {
			t.Fatalf("isVendoredDirectory(%q) = %t, want %t", test.name, got, test.vendored)
		}
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

func newRepositoryFixture() fstest.MapFS {
	return fstest.MapFS{
		agentsFile:          &fstest.MapFile{Data: []byte("short")},
		claudeFile:          &fstest.MapFile{Data: []byte("short")},
		governanceDirectory: &fstest.MapFile{Mode: fs.ModeDir},
	}
}

func repositoryWith(paths ...string) fstest.MapFS {
	fileSystem := fstest.MapFS{
		governanceDirectory: &fstest.MapFile{Mode: fs.ModeDir},
	}
	for _, filePath := range paths {
		putFile(fileSystem, filePath, "short")
	}
	return fileSystem
}

func putFile(fileSystem fstest.MapFS, filePath, contents string) {
	fileSystem[filePath] = &fstest.MapFile{Data: []byte(contents)}
}

func words(count int) string {
	return strings.TrimSpace(strings.Repeat("word ", count))
}
