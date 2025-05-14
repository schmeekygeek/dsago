package quicksort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}
	QuickSort(arr)
	if !reflect.DeepEqual(expected, arr) {
		t.Fatalf("Your array ain't sorted biatch")
	}
}
