package bdd

import (
	"context"
	"io"
	"testing"

	"github.com/cucumber/godog"
)

// Suite executes one canonical corpus through one adapter factory.
type Suite struct {
	Name     string
	Catalog  Catalog
	Factory  DriverFactory
	Registry Registry
	Output   io.Writer
}

// Run executes every recursively discovered scenario serially and strictly.
func (suite Suite) Run(t *testing.T) int {
	t.Helper()
	output := suite.Output
	if output == nil {
		output = io.Discard
	}

	return godog.TestSuite{
		Name: suite.Name,
		ScenarioInitializer: func(scenarioContext *godog.ScenarioContext) {
			var state *State
			scenarioContext.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
				state = NewState(suite.Factory)

				return ctx, nil
			})
			suite.Registry.Register(scenarioContext, &state)
		},
		Options: &godog.Options{
			Format:      "progress",
			NoColors:    true,
			Output:      output,
			Paths:       suite.Catalog.Paths(),
			Strict:      true,
			Concurrency: 1,
		},
	}.Run()
}
