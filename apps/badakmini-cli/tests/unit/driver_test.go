package unit_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/cli"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

// driver executes command orchestration with injected collaborators only. The
// canonical scenarios added in Phase 2 will teach Prepare which deterministic
// fixture each Given step requests.
type driver struct {
	runtime cli.Runtime
	stdout  bytes.Buffer
	stderr  bytes.Buffer
	result  bdd.Result
}

func newDriver() *driver {
	driver := &driver{}
	driver.runtime = cli.Runtime{
		Stdin:  bytes.NewReader(nil),
		Stdout: &driver.stdout,
		Stderr: &driver.stderr,
		FindRepositoryRoot: func() (string, error) {
			return "repository", nil
		},
		CheckGovernance:    func(string) ([]governance.Finding, error) { return nil, nil },
		CheckMarkdownLinks: func(string) ([]markdownlinks.Finding, error) { return nil, nil },
		ListStagedPaths:    func(string) ([]string, error) { return nil, nil },
		CheckParity:        func(string) ([]parity.Finding, error) { return nil, nil },
	}
	return driver
}

func (driver *driver) Prepare(_ context.Context, _ string) error {
	driver.stdout.Reset()
	driver.stderr.Reset()
	driver.result = bdd.Result{}
	return nil
}

func (driver *driver) Invoke(ctx context.Context, args []string) error {
	driver.result.ExitCode = cli.Run(ctx, driver.runtime, args)
	driver.result.Stdout = driver.stdout.String()
	driver.result.Stderr = driver.stderr.String()
	return nil
}

func (driver *driver) Result() bdd.Result {
	return driver.result
}

func TestUnitDriverRunsWithInjectedCollaborators(t *testing.T) {
	adapter := newDriver()
	if err := adapter.Prepare(context.Background(), "foundation"); err != nil {
		t.Fatalf("prepare unit adapter: %v", err)
	}
	if err := adapter.Invoke(context.Background(), []string{"--help"}); err != nil {
		t.Fatalf("invoke unit adapter: %v", err)
	}
	result := adapter.Result()
	if result.ExitCode != 0 || !strings.Contains(result.Stdout, "Usage:") || result.Stderr != "" {
		t.Fatalf("unexpected unit result: %#v", result)
	}
}
