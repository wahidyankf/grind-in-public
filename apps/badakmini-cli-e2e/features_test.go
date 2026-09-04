package e2e_test

import (
	"bytes"
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cucumber/godog"
)

func TestE2EFeatures(t *testing.T) {
	behaviourRoot, err := canonicalBehaviourRoot()
	if err != nil {
		t.Fatalf("locate canonical behaviour: %v", err)
	}
	var output bytes.Buffer
	suite := godog.TestSuite{
		Name: "badakmini-e2e",
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
				return contextWithState(ctx, &scenarioState{driver: newProcessDriver(t)}), nil
			})
			InitializeScenario(sc)
		},
		Options: &godog.Options{
			Format:      "progress",
			NoColors:    true,
			Output:      &output,
			Paths:       []string{behaviourRoot},
			Strict:      true,
			Tags:        "~@e2e-exempt",
			Concurrency: 1,
		},
	}
	if status := suite.Run(); status != 0 {
		t.Fatalf("E2E behaviour suite failed with status %d:\n%s", status, output.String())
	}
}

func canonicalBehaviourRoot() (string, error) {
	_, sourcePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("locate E2E feature suite source")
	}
	return filepath.Clean(filepath.Join(
		filepath.Dir(sourcePath),
		"..", "..",
		"specs", "apps", "badakmini-cli", "behaviours",
	)), nil
}
