package linkedlist

import "fmt"

type Node struct {
  value     int
  next      *Node
  prev      *Node
}

type SinglyLinkedList struct {
  head    Node
  size    int
}

func (list *SinglyLinkedList) AddToLast(value int) {
  node := new(Node)
  if &list.head == nil {
    list.head = *node
  }
  node.value = value
  head := &list.head
  for head.next != nil {
    head = head.next
  }
  head.next = node
  list.size++
}

func (list *SinglyLinkedList) SumList() int {
  head := &list.head
  sum := 0
  for {
    if head == nil {
      break
    }
    sum+=head.value
    head = head.next
  }
  return sum
}

func (list *SinglyLinkedList) PrintList() {
  head := &list.head
  for {
    if head == nil {
      break
    }
    fmt.Printf("%s ", fmt.Sprint(head.value))
    head = head.next
  }
  fmt.Println()
}
