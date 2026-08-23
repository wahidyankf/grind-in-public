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
	behaviorRoot, err := canonicalBehaviorRoot()
	if err != nil {
		t.Fatalf("locate canonical behavior: %v", err)
	}
	var output bytes.Buffer
	suite := godog.TestSuite{
		Name: "badakmini-e2e",
		ScenarioInitializer: func(scenarioContext *godog.ScenarioContext) {
			scenarioContext.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
				return contextWithState(ctx, &scenarioState{driver: newProcessDriver(t)}), nil
			})
			InitializeScenario(scenarioContext)
		},
		Options: &godog.Options{
			Format:      "progress",
			NoColors:    true,
			Output:      &output,
			Paths:       []string{behaviorRoot},
			Strict:      true,
			Concurrency: 1,
		},
	}
	if status := suite.Run(); status != 0 {
		t.Fatalf("E2E behavior suite failed with status %d:\n%s", status, output.String())
	}
}

func canonicalBehaviorRoot() (string, error) {
	_, sourcePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("locate E2E feature suite source")
	}
	return filepath.Clean(filepath.Join(
		filepath.Dir(sourcePath),
		"..", "..", "..",
		"specs", "apps", "badakmini-cli", "behavior",
	)), nil
}
