package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
  expected := []int{1}
  got := topKFrequent([]int{1}, 1)

  if !reflect.DeepEqual(expected, got) {
    t.Fatalf(
      "Failed TopKFrequentElements, expected: %v, got: %v",
      expected,
      got,
    )
  }
}

func TestTopKFrequentElementsTwo(t *testing.T) {
  expected := []int{1, 2}
  got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)

  if !reflect.DeepEqual(expected, got) {
    t.Fatalf(
      "Failed TopKFrequentElements, expected: %v, got: %v",
      expected,
      got,
    )
  }
}
