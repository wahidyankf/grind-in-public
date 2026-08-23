// Package governance validates the repository's concise-governance policy.
package governance

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

const (
	// MaxWords is the hard word limit for each governed Markdown file.
	MaxWords = 500

	agentsFile          = "AGENTS.md"
	claudeFile          = "CLAUDE.md"
	governanceDirectory = "repo-governance"
	readmeFile          = "README.md"
)

// harnessDirectories hold each harness's project configuration, including the
// shared .agents directory more than one harness reads. Their READMEs are
// indexes, so they share the concise-guidance limit, while the agent, skill,
// and command definitions beside them are prompts and stay unmeasured.
var harnessDirectories = []string{".agents", ".claude", ".codex", ".opencode"}

// instructionFiles are the root agent instruction files. They share one limit
// because each must stay equally concise for the harness that reads it.
var instructionFiles = []string{agentsFile, claudeFile}

// Finding describes one governance document that exceeds MaxWords.
type Finding struct {
	Path      string
	WordCount int
}

// CheckFS validates every root instruction file and recursive governance
// document supplied by the production filesystem adapter.
func CheckFS(fileSystem fs.FS) ([]Finding, error) {
	if err := requireGovernanceStructure(fileSystem); err != nil {
		return nil, err
	}

	findings, err := checkInstructionFiles(fileSystem)
	if err != nil {
		return nil, err
	}

	governanceFindings, err := checkGovernanceDocuments(fileSystem)
	if err != nil {
		return nil, err
	}
	findings = append(findings, governanceFindings...)

	harnessFindings, err := checkHarnessReadmes(fileSystem)
	if err != nil {
		return nil, err
	}
	findings = append(findings, harnessFindings...)

	return findings, nil
}

// requireGovernanceStructure distinguishes a missing policy surface from an
// empty, valid one before any word-count work begins.
func requireGovernanceStructure(fileSystem fs.FS) error {
	for _, instructionFile := range instructionFiles {
		if err := requireFile(fileSystem, instructionFile); err != nil {
			return err
		}
	}

	return requireDirectory(fileSystem, governanceDirectory)
}

func checkInstructionFiles(fileSystem fs.FS) ([]Finding, error) {
	var findings []Finding
	for _, instructionFile := range instructionFiles {
		fileFindings, err := checkFile(fileSystem, instructionFile)
		if err != nil {
			return nil, err
		}
		findings = append(findings, fileFindings...)
	}

	return findings, nil
}

func checkGovernanceDocuments(fileSystem fs.FS) ([]Finding, error) {
	var findings []Finding
	err := fs.WalkDir(fileSystem, governanceDirectory, func(path string, entry fs.DirEntry, walkErr error) error {
		return visitGovernanceDocument(fileSystem, path, entry, walkErr, &findings)
	})
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", governanceDirectory, err)
	}

	return findings, nil
}

// visitGovernanceDocument limits the recursive scan to authored Markdown.
func visitGovernanceDocument(
	fileSystem fs.FS,
	filePath string,
	entry fs.DirEntry,
	walkErr error,
	findings *[]Finding,
) error {
	if walkErr != nil {
		return walkErr
	}
	if entry.IsDir() || path.Ext(entry.Name()) != ".md" {
		return nil
	}

	fileFindings, err := checkFile(fileSystem, filePath)
	if err != nil {
		return err
	}
	*findings = append(*findings, fileFindings...)

	return nil
}

// checkHarnessReadmes measures the README index in every harness directory. A
// harness that this repository does not configure is absent rather than empty,
// so a missing directory is skipped instead of failing the run.
func checkHarnessReadmes(fileSystem fs.FS) ([]Finding, error) {
	var findings []Finding

	for _, harnessDirectory := range harnessDirectories {
		harnessFindings, err := checkHarnessDirectory(fileSystem, harnessDirectory)
		if err != nil {
			return nil, err
		}
		findings = append(findings, harnessFindings...)
	}

	return findings, nil
}

func checkHarnessDirectory(fileSystem fs.FS, harnessDirectory string) ([]Finding, error) {
	if _, err := fs.Stat(fileSystem, harnessDirectory); err != nil {
		if errorsIsNotExist(err) {
			return nil, nil
		}

		return nil, fmt.Errorf("inspect %s: %w", harnessDirectory, err)
	}

	var findings []Finding
	err := fs.WalkDir(fileSystem, harnessDirectory, func(filePath string, entry fs.DirEntry, walkErr error) error {
		return visitHarnessReadme(fileSystem, harnessDirectory, filePath, entry, walkErr, &findings)
	})
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", harnessDirectory, err)
	}

	return findings, nil
}

func visitHarnessReadme(
	fileSystem fs.FS,
	harnessPath, filePath string,
	entry fs.DirEntry,
	walkErr error,
	findings *[]Finding,
) error {
	if walkErr != nil {
		return walkErr
	}
	if entry.IsDir() {
		// Harness-local dependency and cache directories are not authored indexes.
		if filePath != harnessPath && isVendoredDirectory(entry.Name()) {
			return fs.SkipDir
		}

		return nil
	}
	if entry.Name() != readmeFile {
		return nil
	}

	fileFindings, err := checkFile(fileSystem, filePath)
	if err != nil {
		return err
	}
	*findings = append(*findings, fileFindings...)

	return nil
}

// isVendoredDirectory reports whether a directory inside a harness holds
// installed or generated content rather than authored configuration. Hidden
// directories are treated the same way because tools put caches there.
func isVendoredDirectory(name string) bool {
	return name == "node_modules" || strings.HasPrefix(name, ".")
}

func requireFile(fileSystem fs.FS, displayPath string) error {
	info, err := fs.Stat(fileSystem, displayPath)
	if err != nil {
		if errorsIsNotExist(err) {
			return fmt.Errorf("required file not found: %s", displayPath)
		}
		return fmt.Errorf("inspect %s: %w", displayPath, err)
	}
	if !info.Mode().IsRegular() {
		// A directory with the expected name must not silently pass as a document.
		return fmt.Errorf("required file is not regular: %s", displayPath)
	}
	return nil
}

func requireDirectory(fileSystem fs.FS, displayPath string) error {
	info, err := fs.Stat(fileSystem, displayPath)
	if err != nil {
		if errorsIsNotExist(err) {
			return fmt.Errorf("required directory not found: %s", displayPath)
		}
		return fmt.Errorf("inspect %s: %w", displayPath, err)
	}
	if !info.IsDir() {
		// Likewise, a file cannot serve as the recursive governance namespace.
		return fmt.Errorf("required path is not a directory: %s", displayPath)
	}
	return nil
}

func checkFile(fileSystem fs.FS, relativePath string) ([]Finding, error) {
	contents, err := fs.ReadFile(fileSystem, relativePath)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", relativePath, err)
	}

	// Fields mirrors the previous shell check's whitespace-based definition of a
	// word and treats Markdown syntax as part of the concise-document budget.
	wordCount := len(strings.Fields(string(contents)))
	if wordCount <= MaxWords {
		return nil, nil
	}

	return []Finding{{Path: relativePath, WordCount: wordCount}}, nil
}

func errorsIsNotExist(err error) bool {
	return errors.Is(err, fs.ErrNotExist)
}
