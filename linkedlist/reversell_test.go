package linkedlist

import "testing"

func TestReverseLL(t *testing.T) {
	list := new(SinglyLinkedList)
	list.AddToLast(1)
	list.AddToLast(2)
	list.AddToLast(3)
	list.AddToLast(4)
	list.PrintList()
	list.reverse()
	list.PrintList()
}
