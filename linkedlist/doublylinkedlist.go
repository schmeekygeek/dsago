package linkedlist

import (
	"errors"
	"fmt"
	"log"
)

type DoublyLinkedList struct {
  size    int
  head    *Node
  tail    *Node
}

func (d *DoublyLinkedList) Prepend(item int) {

  d.size++
  node := Node{
  	value: item,
    prev: nil,
    next: nil,
  }

  if d.head == nil {
    d.head = &node
    d.tail = &node
    return
  }

  node.next = d.head
  node.prev = nil
  d.head.prev = &node
  d.head = &node
}

func (d *DoublyLinkedList) Append(item int) {
  d.size++
  node := Node{
  	value: item,
    prev: nil,
    next: nil,
  }

  if d.head == nil {
    d.head = &node
    d.tail = &node
    return
  }
  node.prev = d.tail
  d.tail.next = &node
  d.tail = &node
}

func (d *DoublyLinkedList) InsertAt(item, idx int) {
  node := Node{
  	value: item,
  	next:  nil,
  	prev:  nil,
  }

  if idx > d.size {
    log.Fatal("Index can't be greater than size")
    return
  } else if d.head == nil {
    d.head = &node
    return
  } else if idx == 0 {
    d.Prepend(item)
    return
  } else if idx == d.size {
    d.Append(item)
    return
  }

  err, curr := getAt(idx, d)
  if err != nil {
    log.Fatal(err)
  }
  node.next = curr.next
  curr.prev = &node
  curr.next = &node
  node.prev = curr
  d.size++
}

func (d *DoublyLinkedList) RemoveAt(idx int) ( error, int ) {
  err, node := getAt(idx, d)
  if err != nil {
    log.Fatal(err)
    return err, node.value
  }
  if node == d.head {
    d.head = d.head.next
  } else if node == d.tail {
    d.tail.prev.next = nil
  } else {
    node.prev.next = node.next
  }
  d.size--
  return err, node.value
}

func (d *DoublyLinkedList) Get(idx int) ( error, int ) {
  err, res := getAt(idx, d)
  return err, res.value
}


func getAt(idx int, d *DoublyLinkedList ) ( error, *Node ) {
  curr := d.head
  if idx >= d.size {
    return errors.New("Index out of bounds"), nil
  }
  for i := 0; i < idx; i++ {
    curr = curr.next
  }
  return nil, curr
}

func (d *DoublyLinkedList) PrintDoubly() {
  head := d.head
  for head != nil {
    fmt.Print(head.value, " ")
    head = head.next
  }
  fmt.Println()
}
