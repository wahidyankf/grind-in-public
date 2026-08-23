package bdd

import (
	"strings"
	"testing"
)

func TestCanonicalBindingsMatchTheRecursiveCorpus(t *testing.T) {
	catalog, err := CanonicalCatalog()
	if err != nil {
		t.Fatalf("discover canonical behavior: %v", err)
	}
	if findings := CanonicalRegistry().Validate(catalog.Steps()); len(findings) != 0 {
		t.Fatalf("validate canonical bindings: %v", findings)
	}
}

func TestBindingComplianceRejectsUndefinedStep(t *testing.T) {
	findings := NewRegistry().Validate([]Step{{Kind: WhenStep, Text: "an undefined action"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected one undefined-step finding, got %v", findings)
	}
}

func TestBindingComplianceRejectsAmbiguousStep(t *testing.T) {
	registry := NewRegistry(
		Definition{Kind: ThenStep, Expression: "the result is (.*)"},
		Definition{Kind: ThenStep, Expression: "the result is good"},
	)
	findings := registry.Validate([]Step{{Kind: ThenStep, Text: "the result is good"}})
	if len(findings) == 0 || !strings.Contains(findings[0].Error(), "ambiguous") {
		t.Fatalf("expected an ambiguous-step finding, got %v", findings)
	}
}

func TestBindingComplianceRejectsUnusedBinding(t *testing.T) {
	registry := NewRegistry(
		Definition{Kind: GivenStep, Expression: "a used fixture"},
		Definition{Kind: GivenStep, Expression: "an unused fixture"},
	)
	findings := registry.Validate([]Step{{Kind: GivenStep, Text: "a used fixture"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "unused") {
		t.Fatalf("expected one unused-binding finding, got %v", findings)
	}
}

func TestBindingComplianceAddedBindingFixture(t *testing.T) {
	registry := NewRegistry(
		Definition{Kind: GivenStep, Expression: "a used fixture"},
		Definition{Kind: GivenStep, Expression: "an added fixture"},
	)
	findings := registry.Validate([]Step{{Kind: GivenStep, Text: "a used fixture"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "unused") {
		t.Fatalf("expected added binding to be unused, got %v", findings)
	}
}

func TestBindingComplianceEditedBindingFixture(t *testing.T) {
	registry := NewRegistry(Definition{Kind: WhenStep, Expression: "the edited action"})
	findings := registry.Validate([]Step{{Kind: WhenStep, Text: "the original action"}})
	if len(findings) != 2 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected edited binding mismatch, got %v", findings)
	}
}

func TestBindingComplianceRenamedBindingFixture(t *testing.T) {
	registry := NewRegistry(Definition{Kind: ThenStep, Expression: "the renamed result"})
	findings := registry.Validate([]Step{{Kind: ThenStep, Text: "the old result"}})
	if len(findings) != 2 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected renamed binding mismatch, got %v", findings)
	}
}

func TestBindingComplianceDeletedBindingFixture(t *testing.T) {
	findings := NewRegistry().Validate([]Step{{Kind: GivenStep, Text: "a deleted fixture"}})
	if len(findings) != 1 || !strings.Contains(findings[0].Error(), "undefined") {
		t.Fatalf("expected deleted binding mismatch, got %v", findings)
	}
}
