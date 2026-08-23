package bdd

import (
	"reflect"
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
