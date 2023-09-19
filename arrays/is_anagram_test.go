package arrays

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
  expected := false
  got := IsAnagram("rat", "car")
  if expected != got {
    t.Fatalf("Test1 failed")
  }
}

func TestIsAnagramTwo(t *testing.T) {
  expected := true
  got := IsAnagram("anagram", "nagaram")
  if expected != got {
    t.Fatalf("Test2 failed")
  }
}

func TestIsAnagramThree(t *testing.T) {
  expected := true
  got := IsAnagram("nojma", "manoj")
  if expected != got {
    t.Fatalf("Test2 failed")
  }
}
