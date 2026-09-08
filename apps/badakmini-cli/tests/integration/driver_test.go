package integration_test

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

type behaviourDriver struct {
	testing *testing.T
	root    string
	stdout  bytes.Buffer
	stderr  bytes.Buffer
	runtime cli.Runtime
	result  bdd.Result
}

func newBehaviourDriver(t *testing.T) *behaviourDriver {
	t.Helper()
	driver := &behaviourDriver{testing: t, root: t.TempDir()}
	driver.resetRuntime()
	return driver
}

//nolint:funcorder // Constructor setup stays beside newBehaviourDriver for the fixture lifecycle.
func (driver *behaviourDriver) resetRuntime() {
	driver.stdout.Reset()
	driver.stderr.Reset()
	driver.result = bdd.Result{}
	driver.runtime = cli.Runtime{
		Stdin:              strings.NewReader(""),
		Stdout:             &driver.stdout,
		Stderr:             &driver.stderr,
		FindRepositoryRoot: func() (string, error) { return driver.root, nil },
		ListStagedPaths:    realStagedPaths,
	}
}

func (driver *behaviourDriver) Prepare(_ context.Context, fixture string) error {
	driver.resetRuntime()
	switch fixture {
	case "repository-discovery-fails":
		driver.runtime.FindRepositoryRoot = func() (string, error) {
			return "", errors.New("repository discovery failed")
		}
	case "staged-rule-bearing-file":
		driver.stageFile("repo-governance/development/testing-policy.md")
	case "ordinary-staged-file":
		driver.stageFile("README.md")
	case "harness-instruction-pre-edit":
		driver.runtime.Stdin = strings.NewReader(`{"tool_input":{"file_path":"AGENTS.md"}}`)
	default:
		return errors.New("unsupported integration fixture: " + fixture)
	}
	return nil
}

func (driver *behaviourDriver) Invoke(ctx context.Context, arguments []string) error {
	driver.result.ExitCode = cli.Run(ctx, driver.runtime, arguments)
	driver.result.Stdout = driver.stdout.String()
	driver.result.Stderr = driver.stderr.String()
	return nil
}

func (driver *behaviourDriver) Result() bdd.Result {
	return driver.result
}

func (driver *behaviourDriver) stageFile(path string) {
	runRuleChangeGit(driver.testing, driver.root, "init", "--quiet")
	writeRuleChangeFile(driver.testing, driver.root, path, "fixture")
	runRuleChangeGit(driver.testing, driver.root, "add", "--", path)
}
