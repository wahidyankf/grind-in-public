package bdd

import (
	"reflect"
	"testing"
)

func TestDriverComplianceRejectsIncompleteAdapter(t *testing.T) {
	missing := MissingDriverMethods(reflect.TypeFor[*incompleteDriverFixture]())
	if !reflect.DeepEqual(missing, []string{"Result"}) {
		t.Fatalf("expected missing Result operation, got %v", missing)
	}
}

func TestDriverComplianceAcceptsCompleteAdapter(t *testing.T) {
	if missing := MissingDriverMethods(reflect.TypeFor[*fixtureDriver]()); len(missing) != 0 {
		t.Fatalf("expected complete adapter, missing %v", missing)
	}
}
