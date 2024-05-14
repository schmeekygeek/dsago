package arraysandhashing

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }

  freq1 := make(map[string]int)
  freq2 := make(map[string]int)
  for i := 0; i < len(s); i++ {
    freq1[string(s[i])] = freq1[string(s[i])] + 1
    freq2[string(t[i])] = freq2[string(t[i])] + 1
  }
  fmt.Println(freq1)
  fmt.Println(freq2)

  return reflect.DeepEqual(freq1, freq2) 
}
