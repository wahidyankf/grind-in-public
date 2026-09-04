package bdd

import (
	"strings"
	"testing"
	"testing/fstest"
)

func TestFeatureComplianceRejectsMalformedFeature(t *testing.T) {
	_, err := discoverSource("Scenario: Orphan\n  Given a fixture\n  When it runs\n  Then it succeeds\n")
	if err == nil || !strings.Contains(err.Error(), "parse behaviour features with Godog") {
		t.Fatalf("expected malformed feature rejection, got %v", err)
	}
}

func TestFeatureComplianceRejectsEmptyFeature(t *testing.T) {
	_, err := discoverSource("Feature: Empty\n")
	if err == nil || !strings.Contains(err.Error(), "at least one scenario") {
		t.Fatalf("expected empty feature rejection, got %v", err)
	}
}

func TestFeatureComplianceRequiresOnePrimaryGivenWhenThen(t *testing.T) {
	testCases := map[string]string{
		"missing Given": "Feature: Invalid\nScenario: Missing\nWhen it runs\nThen it succeeds\n",
		"missing When":  "Feature: Invalid\nScenario: Missing\nGiven a fixture\nThen it succeeds\n",
		"missing Then":  "Feature: Invalid\nScenario: Missing\nGiven a fixture\nWhen it runs\n",
		"duplicate Given": "Feature: Invalid\nScenario: Duplicate\nGiven a fixture\n" +
			"Given another fixture\nWhen it runs\nThen it succeeds\n",
		"duplicate When": "Feature: Invalid\nScenario: Duplicate\nGiven a fixture\n" +
			"When it runs\nWhen it runs again\nThen it succeeds\n",
		"duplicate Then": "Feature: Invalid\nScenario: Duplicate\nGiven a fixture\n" +
			"When it runs\nThen it succeeds\nThen it succeeds again\n",
	}
	for name, source := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := discoverSource(source)
			if err == nil || !strings.Contains(err.Error(), "requires exactly one primary") {
				t.Fatalf("expected primary-step rejection, got %v", err)
			}
		})
	}
}

func TestFeatureComplianceCountsExpandedScenarioOutlineRows(t *testing.T) {
	source := "Feature: Outline\nScenario Outline: Expand\nGiven a <fixture>\n" +
		"When it runs\nThen it succeeds\nExamples:\n| fixture |\n| first |\n| second |\n"
	catalog, err := discoverSource(source)
	if err != nil {
		t.Fatalf("discover scenario outline: %v", err)
	}
	if catalog.Features[0].ScenarioCount != 2 {
		t.Fatalf("expected two expanded scenarios, got %d", catalog.Features[0].ScenarioCount)
	}
	if len(catalog.Features[0].Steps) != 6 {
		t.Fatalf("expected six expanded steps, got %d", len(catalog.Features[0].Steps))
	}
	if catalog.Features[0].Steps[0].Text != "a first" || catalog.Features[0].Steps[3].Text != "a second" {
		t.Fatalf("expected Godog-expanded outline steps, got %#v", catalog.Features[0].Steps)
	}
}

func TestFeatureComplianceUsesGodogParserForMalformedExamples(t *testing.T) {
	source := "Feature: Outline\nScenario Outline: Expand\nGiven a <fixture>\n" +
		"When it runs\nThen it succeeds\nExamples:\n| fixture |\n| first | extra |\n"
	if _, err := discoverSource(source); err == nil {
		t.Fatal("expected Godog to reject an examples row with the wrong cell count")
	}
}

func discoverSource(source string) (Catalog, error) {
	return DiscoverFS(fstest.MapFS{
		"behaviour.feature": &fstest.MapFile{Data: []byte(source)},
	}, ".")
}
