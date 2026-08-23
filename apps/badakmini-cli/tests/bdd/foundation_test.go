package bdd

import (
	"context"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/cucumber/godog"
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
	initializeScenario := func(sc *godog.ScenarioContext) {
		sc.Given(`^a foundation fixture$`, prepareFixture("foundation"))
		sc.When(`^the fixture runs$`, invokeCommand())
		sc.Then(`^the fixture succeeds$`, func(context.Context) error { return nil })
	}

	suite := Suite{
		Name:                "foundation",
		Catalog:             catalog,
		Factory:             func() Driver { return &fixtureDriver{} },
		ScenarioInitializer: initializeScenario,
	}
	if status := suite.Run(t); status != 0 {
		t.Fatalf("expected fixture suite success, got status %d", status)
	}
}

func TestFoundationRegistersBindingsByGherkinKeyword(t *testing.T) {
	feature := "Feature: Keyword bindings\nScenario: Shared text\n" +
		"Given the shared phrase\nWhen the fixture runs\nThen the shared phrase\n"
	filesystem := fstest.MapFS{"keyword.feature": &fstest.MapFile{Data: []byte(feature)}}
	catalog, err := DiscoverFS(filesystem, ".")
	if err != nil {
		t.Fatalf("discover keyword feature: %v", err)
	}
	initializeScenario := func(sc *godog.ScenarioContext) {
		sc.Given(`^the shared phrase$`, func(context.Context) error { return nil })
		sc.When(`^the fixture runs$`, func(context.Context) error { return nil })
		sc.Then(`^the shared phrase$`, func(context.Context) error { return nil })
	}

	suite := Suite{
		Name:                "keyword-bindings",
		Catalog:             catalog,
		Factory:             func() Driver { return &fixtureDriver{} },
		ScenarioInitializer: initializeScenario,
		FS:                  filesystem,
	}
	if status := suite.Run(t); status != 0 {
		t.Fatalf("expected keyword-specific Godog bindings to pass, got status %d", status)
	}
}

func TestStateCreatesFreshDriver(t *testing.T) {
	first := NewState(func() Driver { return &fixtureDriver{} })
	second := NewState(func() Driver { return &fixtureDriver{} })
	if first.Driver() == second.Driver() {
		t.Fatal("expected scenario state to own a fresh driver")
	}
}
