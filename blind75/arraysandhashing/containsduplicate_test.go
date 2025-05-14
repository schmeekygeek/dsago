package arraysandhashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	expected := true
	got := containsDuplicate(nums)
	if expected != got {
		t.Fatalf("Failed ContainsDuplicate, expected: %t, got: %t", expected, got)
	}
}

func TestContainsDuplicate2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := false
	got := containsDuplicate(nums)
	if expected != got {
		t.Fatalf("Failed ContainsDuplicate, expected: %t, got: %t", expected, got)
	}
}
