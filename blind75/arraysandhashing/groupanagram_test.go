package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
  expected := [][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}}
  got := groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})

  if reflect.DeepEqual(expected, got) {
    t.Fatalf("Failed TestGroupAnagram, expected: %v got: %v", expected, got)
  }
}

func TestGroupAnagramTwo(t *testing.T) {
  expected := [][]string{{"bat"}}
  got := groupAnagrams([]string{"bat"})

  if !reflect.DeepEqual(expected, got) {
    t.Fatalf("Failed TestGroupAnagram, expected: %v got: %v", expected, got)
  }
}
