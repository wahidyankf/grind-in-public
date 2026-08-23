package bdd

import (
	"context"
	"path/filepath"
	"testing"
)

type fixtureDriver struct {
	prepared bool
	result   Result
}

// incompleteDriverFixture deliberately omits Result for driver-compliance fixtures.
type incompleteDriverFixture struct{}

func (*incompleteDriverFixture) Prepare(context.Context, string) error { return nil }

func (*incompleteDriverFixture) Invoke(context.Context, []string) error { return nil }

func (driver *fixtureDriver) Prepare(context.Context, string) error {
	driver.prepared = true

	return nil
}

func (driver *fixtureDriver) Invoke(context.Context, []string) error {
	driver.result = Result{ExitCode: 0, Stdout: "fixture passed"}

	return nil
}

func (driver *fixtureDriver) Result() Result {
	return driver.result
}

func TestFoundationExecutesFixtureFeature(t *testing.T) {
	catalog, err := Discover(filepath.Join("testdata", "foundation"))
	if err != nil {
		t.Fatalf("discover fixture feature: %v", err)
	}
	registry := NewRegistry(
		Definition{Kind: GivenStep, Expression: "a foundation fixture", Handler: func(state *State) any {
			return func() error { return state.Driver().Prepare(context.Background(), "foundation") }
		}},
		Definition{Kind: WhenStep, Expression: "the fixture runs", Handler: func(state *State) any {
			return func() error { return state.Driver().Invoke(context.Background(), nil) }
		}},
		Definition{Kind: ThenStep, Expression: "the fixture succeeds", Handler: func(_ *State) any {
			return func() error { return nil }
		}},
	)
	if findings := registry.Validate(catalog.Steps()); len(findings) != 0 {
		t.Fatalf("validate fixture bindings: %v", findings)
	}

	suite := Suite{
		Name:     "foundation",
		Catalog:  catalog,
		Factory:  func() Driver { return &fixtureDriver{} },
		Registry: registry,
	}
	if status := suite.Run(t); status != 0 {
		t.Fatalf("expected fixture suite success, got status %d", status)
	}
}

func TestStateCreatesFreshDriver(t *testing.T) {
	first := NewState(func() Driver { return &fixtureDriver{} })
	second := NewState(func() Driver { return &fixtureDriver{} })
	if first.Driver() == second.Driver() {
		t.Fatal("expected scenario state to own a fresh driver")
	}
}
