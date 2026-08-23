package bdd

import "testing"

func TestUnitBoundaryPolicy(t *testing.T) {
	findings, err := unitBoundaryFindings()
	if err != nil {
		t.Fatalf("inspect unit test boundaries: %v", err)
	}
	if len(findings) != 0 {
		for _, finding := range findings {
			t.Error(finding)
		}
	}
}

func TestIntegrationBoundaryPolicy(t *testing.T) {
	findings, err := integrationBoundaryFindings()
	if err != nil {
		t.Fatalf("inspect integration test boundaries: %v", err)
	}
	if len(findings) != 0 {
		for _, finding := range findings {
			t.Error(finding)
		}
	}
}

func TestE2EBoundaryPolicy(t *testing.T) {
	findings, err := e2eBoundaryFindings()
	if err != nil {
		t.Fatalf("inspect E2E test boundary: %v", err)
	}
	if len(findings) != 0 {
		for _, finding := range findings {
			t.Error(finding)
		}
	}
}
