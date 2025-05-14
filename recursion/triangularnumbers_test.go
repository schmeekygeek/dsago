package recursion

import (
	"testing"
)

func TestTriangularNumbers(t *testing.T) {
	// expected
	exp := 28

	// got
	got := nthtriangular(7)

	if exp != got {
		t.Fatal("TestTriangularNumbers failed")
	}
}
