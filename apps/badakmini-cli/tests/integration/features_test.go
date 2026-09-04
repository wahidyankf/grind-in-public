package integration_test

import (
	"testing"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/tests/bdd"
)

func TestIntegrationFeatures(t *testing.T) {
	catalog, err := bdd.CanonicalCatalog()
	if err != nil {
		t.Fatalf("discover canonical behaviour: %v", err)
	}

	suite := bdd.Suite{
		Name:    "badakmini-integration",
		Catalog: catalog,
		Factory: func() bdd.Driver { return newBehaviourDriver(t) },
		Tags:    "~@integration-exempt",
	}
	if status := suite.Run(t); status != 0 {
		t.Fatalf("integration behaviour suite failed with status %d", status)
	}
}
