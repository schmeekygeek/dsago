package linkedlist

import (
	"testing"
)

func TestDoubly(t *testing.T) {
	dl := DoublyLinkedList{
		size: 0,
		head: nil,
	}
	dl.Append(1)
	dl.Append(2)
	dl.Append(3)
	dl.InsertAt(1, 0)
	err, val := dl.Get(1)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(val)
	}

	t.Log(dl.size)
	dl.PrintDoubly()
}
