package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{1, 0}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("TwoSum failed: expected: %v, got: %v", expected, got)
	}
}
