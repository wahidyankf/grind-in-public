package bdd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"
	"testing/fstest"
)

const validFeatureSource = "Feature: Valid\nScenario: Valid\nGiven a fixture\nWhen it runs\nThen it succeeds\n"

func TestAdapterParityDiscoversTheExactRecursiveCorpus(t *testing.T) {
	filesystem := fstest.MapFS{
		"z.feature":        &fstest.MapFile{Data: []byte(validFeatureSource)},
		"nested/a.feature": &fstest.MapFile{Data: []byte(validFeatureSource)},
		"nested/notes.md":  &fstest.MapFile{Data: []byte("ignored")},
	}
	catalog, err := DiscoverFS(filesystem, ".")
	if err != nil {
		t.Fatalf("discover recursive corpus: %v", err)
	}
	expected := []string{"nested/a.feature", "z.feature"}
	if !reflect.DeepEqual(catalog.Paths(), expected) {
		t.Fatalf("expected corpus %v, got %v", expected, catalog.Paths())
	}
}

func TestAdapterParityRejectsCorpusMismatch(t *testing.T) {
	first, err := discoverSource(validFeatureSource)
	if err != nil {
		t.Fatalf("discover first corpus: %v", err)
	}
	second, err := discoverSource(strings.Replace(validFeatureSource, "Scenario: Valid", "Scenario: Different", 1))
	if err != nil {
		t.Fatalf("discover second corpus: %v", err)
	}
	if err := ValidateAdapterParity(map[string]Catalog{"unit": first, "integration": second}); err == nil {
		t.Fatal("expected adapter corpus mismatch")
	}
}

func TestAdapterParityRejectsLayerFilteringTag(t *testing.T) {
	_, err := discoverSource("@unit\n" + validFeatureSource)
	if err == nil || !strings.Contains(err.Error(), "layer-filter") {
		t.Fatalf("expected layer-filter rejection, got %v", err)
	}
}

func TestAdaptersShareCanonicalCatalogAndBindings(t *testing.T) {
	catalog, err := CanonicalCatalog()
	if err != nil {
		t.Fatalf("discover canonical catalog: %v", err)
	}
	root, err := moduleRoot()
	if err != nil {
		t.Fatalf("locate module root: %v", err)
	}
	ownerAdapterSources := map[string]string{
		"unit":        filepath.Join(root, "tests", "unit", "features_test.go"),
		"integration": filepath.Join(root, "tests", "integration", "features_test.go"),
	}
	adapters := make(map[string]Catalog, len(ownerAdapterSources)+1)
	for name, sourcePath := range ownerAdapterSources {
		assertOwnerAdapterSource(t, name, sourcePath)
		adapters[name] = catalog
	}
	e2eSourcePath := filepath.Join(root, "tests", "e2e", "features_test.go")
	assertE2EAdapterSource(t, e2eSourcePath)
	adapters["e2e"] = catalog
	if err := ValidateAdapterParity(adapters); err != nil {
		t.Fatalf("validate adapter catalog parity: %v", err)
	}
	if findings := validateBindings(canonicalBindings(t), catalog.Steps()); len(findings) != 0 {
		t.Fatalf("validate owner bindings: %v", findings)
	}
	if findings := validateBindings(e2eBindings(t), catalog.Steps()); len(findings) != 0 {
		t.Fatalf("validate E2E bindings: %v", findings)
	}
}

func assertOwnerAdapterSource(t *testing.T, name, sourcePath string) {
	t.Helper()
	source := readAdapterSource(t, name, sourcePath)
	if !strings.Contains(source, "bdd.CanonicalCatalog()") {
		t.Errorf("%s adapter must load bdd.CanonicalCatalog", name)
	}
	if strings.Contains(source, "ScenarioInitializer:") {
		t.Errorf("%s adapter must use the shared direct Godog initializer", name)
	}
}

func assertE2EAdapterSource(t *testing.T, sourcePath string) {
	t.Helper()
	source := readAdapterSource(t, "E2E", sourcePath)
	for _, required := range []string{
		`"github.com/cucumber/godog"`,
		"godog.TestSuite{",
		"ScenarioInitializer:",
		"InitializeScenario(sc)",
	} {
		if !strings.Contains(source, required) {
			t.Errorf("E2E adapter must directly own %q", required)
		}
	}
	if strings.Contains(source, "badakmini-cli/tests/bdd") {
		t.Error("E2E adapter must not execute through the owner BDD suite")
	}
}

func readAdapterSource(t *testing.T, name, sourcePath string) string {
	t.Helper()
	// #nosec G304 -- callers pass fixed adapter locations under the checked repository.
	source, err := os.ReadFile(sourcePath)
	if err != nil {
		t.Fatalf("read %s adapter: %v", name, err)
	}
	return string(source)
}

func TestAdapterParityAddedFeatureFixture(t *testing.T) {
	base := fstest.MapFS{"base.feature": &fstest.MapFile{Data: []byte(validFeatureSource)}}
	changed := fstest.MapFS{
		"base.feature":  &fstest.MapFile{Data: []byte(validFeatureSource)},
		"added.feature": &fstest.MapFile{Data: []byte(validFeatureSource)},
	}
	assertCatalogMismatch(t, base, changed)
}

func TestAdapterParityEditedStepFixture(t *testing.T) {
	assertCatalogMismatch(t,
		fstest.MapFS{"behavior.feature": &fstest.MapFile{Data: []byte(validFeatureSource)}},
		fstest.MapFS{
			"behavior.feature": &fstest.MapFile{
				Data: []byte(strings.Replace(validFeatureSource, "it succeeds", "it fails", 1)),
			},
		},
	)
}

func TestAdapterParityRenamedFeatureFixture(t *testing.T) {
	assertCatalogMismatch(t,
		fstest.MapFS{"old.feature": &fstest.MapFile{Data: []byte(validFeatureSource)}},
		fstest.MapFS{"new.feature": &fstest.MapFile{Data: []byte(validFeatureSource)}},
	)
}

func TestAdapterParityNestedFeatureFixture(t *testing.T) {
	filesystem := fstest.MapFS{"nested/deeper/behavior.feature": &fstest.MapFile{Data: []byte(validFeatureSource)}}
	catalog, err := DiscoverFS(filesystem, ".")
	if err != nil {
		t.Fatalf("discover nested feature: %v", err)
	}
	if !reflect.DeepEqual(catalog.Paths(), []string{"nested/deeper/behavior.feature"}) {
		t.Fatalf("expected nested feature in catalog, got %v", catalog.Paths())
	}
}

func TestAdapterParityDeletedFeatureFixture(t *testing.T) {
	assertCatalogMismatch(t,
		fstest.MapFS{
			"first.feature":  &fstest.MapFile{Data: []byte(validFeatureSource)},
			"second.feature": &fstest.MapFile{Data: []byte(validFeatureSource)},
		},
		fstest.MapFS{"first.feature": &fstest.MapFile{Data: []byte(validFeatureSource)}},
	)
}

func TestE2EBindingInputRegression(t *testing.T) {
	inputs := behaviorTargetInputs(t)
	if !slicesContain(inputs, "{workspaceRoot}/apps/badakmini-cli/tests/e2e/**/*") {
		t.Fatalf("behavior target must invalidate for E2E binding changes: %v", inputs)
	}
}

func TestE2EConfigurationInputRegression(t *testing.T) {
	inputs := behaviorTargetInputs(t)
	if !slicesContain(inputs, "default") {
		t.Fatalf("behavior target must invalidate for owner configuration changes: %v", inputs)
	}
}

func assertCatalogMismatch(t *testing.T, firstFiles, secondFiles fstest.MapFS) {
	t.Helper()
	first, err := DiscoverFS(firstFiles, ".")
	if err != nil {
		t.Fatalf("discover first fixture: %v", err)
	}
	second, err := DiscoverFS(secondFiles, ".")
	if err != nil {
		t.Fatalf("discover second fixture: %v", err)
	}
	if err := ValidateAdapterParity(map[string]Catalog{"first": first, "second": second}); err == nil {
		t.Fatal("expected catalog mismatch")
	}
}

func behaviorTargetInputs(t *testing.T) []string {
	t.Helper()
	root, err := moduleRoot()
	if err != nil {
		t.Fatalf("locate module root: %v", err)
	}
	// #nosec G304 -- this is the fixed owner project descriptor adjacent to the test package.
	contents, err := os.ReadFile(filepath.Join(root, "project.json"))
	if err != nil {
		t.Fatalf("read owner project: %v", err)
	}
	var project struct {
		Targets map[string]struct {
			Inputs []string `json:"inputs"`
		} `json:"targets"`
	}
	if err := json.Unmarshal(contents, &project); err != nil {
		t.Fatalf("parse owner project: %v", err)
	}
	target, ok := project.Targets["test:coverage:behavior"]
	if !ok {
		t.Fatal("owner test:coverage:behavior target is missing")
	}
	return target.Inputs
}

func slicesContain(values []string, wanted string) bool {
	return slices.Contains(values, wanted)
}
