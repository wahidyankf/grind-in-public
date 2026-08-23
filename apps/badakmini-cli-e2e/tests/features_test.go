package e2e_test

import (
	"bytes"
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

func TestE2EFeatures(t *testing.T) {
	catalog, err := bdd.CanonicalCatalog()
	if err != nil {
		t.Fatalf("discover canonical behavior: %v", err)
	}
	registry := bdd.CanonicalRegistry()
	if findings := registry.Validate(catalog.Steps()); len(findings) != 0 {
		t.Fatalf("validate canonical bindings: %v", findings)
	}
	var output bytes.Buffer
	suite := bdd.Suite{
		Name:     "badakmini-e2e",
		Catalog:  catalog,
		Factory:  func() bdd.Driver { return newProcessDriver(t) },
		Registry: registry,
		Output:   &output,
	}
	if status := suite.Run(t); status != 0 {
		t.Fatalf("E2E behavior suite failed with status %d:\n%s", status, output.String())
	}
}
