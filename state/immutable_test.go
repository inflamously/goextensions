package state

import (
	"testing"
)

type TestImmutableData struct {
	name  string
	value *TestImmutableData
}

func TestImmutable(t *testing.T) {
	val := "abc"
	immutable := NewImmutable(val)
	if immutable.Get() != "abc" {
		t.Fatalf("Immutable was either not set or is invalid")
	}
}

func TestImmutableWithOpenFields(t *testing.T) {
	val := TestImmutableData{
		name: "x",
		value: &TestImmutableData{
			name:  "y",
			value: nil,
		},
	}

	immutable := NewImmutable(val)
	afterVal := immutable.Get().(TestImmutableData)
	afterVal.name = "abc"
	afterVal.value.name = "def"
	resultImmutable := immutable.Get().(TestImmutableData)
	if resultImmutable.name != "x" {
		t.Fatalf("FATAL: Immutable was changed!")
	}
	if resultImmutable.value.name != "y" {
		t.Fatalf("FATAL: Immutable was changed!")
	}
}
