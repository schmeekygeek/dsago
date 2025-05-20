package twopointers_test

import (
	"dsa/twopointers"
	"reflect"
	"testing"
)

func TestTwoPointers(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12, 0}
	expected := []int{1, 3, 12, 0, 0, 0}
	got := twopointers.MoveZeroes(nums)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("TestTwoPointers failed, expected: %v, got: %v", expected, got)
	}
}
