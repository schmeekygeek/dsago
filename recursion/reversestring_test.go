package recursion

import (
	"testing"
)

func TestReverseString(t *testing.T) {

	// expected
	exp := "dcba"

	// got
	got := reverse("abcd")

	if exp != got {
		t.Fatal("Failed reverse string")
	}
}
