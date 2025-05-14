package arrays_test

import (
	"dsa/arrays"
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	testArray := []int{3, 4, 5, 6, 19, 23, 34, 35, 38, 45, 46, 88}
	target := 38
	expected := 8
	got := arrays.BinarySearch(testArray, target)

	if expected != got {
		log.Println("hi")
		t.Fatalf("TestBinarySearch failed, expected: %v, got: %v", expected, got)
	}
}
