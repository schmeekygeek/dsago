package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
  heap := new(MinHeap)
  heap.Insert(50)
  heap.Insert(71)
  heap.Insert(100)
  heap.Insert(101)
  heap.Insert(80)
  heap.Insert(200)
  heap.printHeap()
  fmt.Println("==========")
  heap.Delete()
  heap.printHeap()
}
