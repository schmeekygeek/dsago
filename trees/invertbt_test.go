package trees

import (
  "reflect"
  "testing"
)

func TestInvertBt(t *testing.T) {
  node := CreateBT()

  // got
  got := make([]int, 0)
  node.DfsInOrder(&got)
  Invert(&node)

  // expected
  expected := []int{5, 23, 4, 7, 18, 3, 21}

  if !reflect.DeepEqual(got, expected) {
    t.Fatal("Invert Binary Tree failed")
  }
}
