package linkedlist

import "fmt"

func (list *SinglyLinkedList) reverse() {
	curr := list.head
	var prev *Node
	for &curr != nil {
		fmt.Println(curr.value)
		next := curr.next
		curr.next = prev
		prev = &curr
		curr = *next
	}
}
