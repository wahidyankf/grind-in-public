package bdd

import (
	"context"
	"io"
	"io/fs"
	"testing"

	"github.com/cucumber/godog"
)

// Suite executes one canonical corpus through one adapter factory.
type Suite struct {
	Name                string
	Catalog             Catalog
	Factory             DriverFactory
	ScenarioInitializer func(*godog.ScenarioContext)
	Output              io.Writer
	FS                  fs.FS
	Tags                string
}

// Run executes every recursively discovered scenario serially and strictly.
func (suite Suite) Run(t *testing.T) int {
	t.Helper()
	output := suite.Output
	if output == nil {
		output = io.Discard
	}
	initializer := suite.ScenarioInitializer
	if initializer == nil {
		initializer = InitializeScenario
	}

	return godog.TestSuite{
		Name: suite.Name,
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			sc.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
				return contextWithState(ctx, NewState(suite.Factory)), nil
			})
			initializer(sc)
		},
		Options: &godog.Options{
			Format:      "progress",
			FS:          suite.FS,
			NoColors:    true,
			Output:      output,
			Paths:       suite.Catalog.Paths(),
			Strict:      true,
			Tags:        suite.Tags,
			Concurrency: 1,
		},
	}.Run()
}
