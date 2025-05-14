package arraysandhashing

import "testing"

func TestIsAnagram(t *testing.T) {
	expected := true
	got := isAnagram("anagram", "nagaram")
	if expected != got {
		t.Fatalf("ValidAnagramOne failed: expected: %t, got: %t", expected, got)
	}
}

func TestIsAnagramThree(t *testing.T) {
	expected := true
	got := isAnagram("aoeu", "ueoa")
	if expected != got {
		t.Fatalf("ValidAnagramTwo failed: expected: %t, got: %t", expected, got)
	}
}

func TestIsAnagramTwo(t *testing.T) {
	expected := false
	got := isAnagram("carnto", "rat")
	if expected != got {
		t.Fatalf("ValidAnagramThree failed: expected: %t, got: %t", expected, got)
	}
}
