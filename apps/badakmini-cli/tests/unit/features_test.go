package unit_test

import (
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

func TestFeatures(t *testing.T) {
	catalog, err := bdd.CanonicalCatalog()
	if err != nil {
		t.Fatalf("discover canonical behavior: %v", err)
	}

	suite := bdd.Suite{
		Name:    "badakmini-unit",
		Catalog: catalog,
		Factory: func() bdd.Driver { return newDriver() },
	}
	if status := suite.Run(t); status != 0 {
		t.Fatalf("unit behavior suite failed with status %d", status)
	}
}
