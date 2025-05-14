package stacks_test

import (
	"dsa/stacks"
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	expected := true
	got := stacks.IsValid("()[]{}")

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("TestStacks failed, expected: %v, got: %v", expected, got)
	}
}
