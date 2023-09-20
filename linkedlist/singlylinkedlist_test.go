package linkedlist

import (
  "testing"
)

func TestLinkedList(t *testing.T) {
  list := new(SinglyLinkedList)
  list.AddToLast(1)
  list.AddToLast(2)
  list.AddToLast(3)
  list.AddToLast(4)
  sum := list.SumList()
  if sum != 10 {
    t.Fatalf("Failed")
  }
}
