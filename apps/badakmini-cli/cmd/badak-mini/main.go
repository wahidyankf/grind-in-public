// Badak Mini is a focused repository-governance command-line tool.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/rulechange"
)

func main() {
	os.Exit(execute(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

func execute(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	return cli.Run(context.Background(), productionRuntime(stdin, stdout, stderr), args)
}

func productionRuntime(stdin io.Reader, stdout, stderr io.Writer) cli.Runtime {
	return cli.Runtime{
		Stdin:              stdin,
		Stdout:             stdout,
		Stderr:             stderr,
		FindRepositoryRoot: findRepositoryRoot,
		CheckGovernance:    checkGovernance,
		CheckMarkdownLinks: checkMarkdownLinks,
		ListStagedPaths:    listStagedPaths,
		CheckParity:        checkParity,
	}
}

func findRepositoryRoot() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", errors.New("run this command from inside a Git repository")
	}
	return strings.TrimSpace(string(output)), nil
}

func checkGovernance(root string) ([]governance.Finding, error) {
	findings, err := governance.CheckFS(os.DirFS(root))
	if err != nil {
		return nil, fmt.Errorf("check governance: %w", err)
	}
	return findings, nil
}

func checkMarkdownLinks(root string) ([]markdownlinks.Finding, error) {
	findings, err := markdownlinks.Check(root, markdownlinks.Runtime{
		ReadFile:     os.ReadFile,
		Stat:         os.Stat,
		EvalSymlinks: filepath.EvalSymlinks,
		TrackedFiles: listTrackedFiles,
	})
	if err != nil {
		return nil, fmt.Errorf("check Markdown links: %w", err)
	}
	return findings, nil
}

func listTrackedFiles(root string) (map[string]struct{}, error) {
	// #nosec G204 -- the executable and arguments are fixed; root is one structured Git argument.
	output, err := exec.Command("git", "-C", root, "ls-files", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("list tracked repository files: %w", err)
	}
	return markdownlinks.ParseTrackedFiles(output), nil
}

func listStagedPaths(root string) ([]string, error) {
	// #nosec G204 -- the executable and arguments are fixed; root is one structured Git argument.
	output, err := exec.Command("git", "-C", root, "diff", "--cached", "--name-only", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("list staged paths: %w", err)
	}
	return rulechange.ParseStagedPaths(output), nil
}

func checkParity(root string) (parity.Report, error) {
	report, err := parity.CheckFS(os.DirFS(root))
	if err != nil {
		return parity.Report{}, fmt.Errorf("check harness parity: %w", err)
	}
	return report, nil
}
