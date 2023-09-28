package heap

import "math"

type MinHeap struct {
  Length  int
  data    []int
}

func (mn *MinHeap) Insert() {
}

func (mn *MinHeap) Delete() {
}

func (mn *MinHeap) heapifyUp(idx int) {
  if idx == 0 {
    return
  }
  parent := mn.parent(idx)
  parentV := mn.data[parent]
  value := mn.data[idx]

  if parent > value {
    mn.data[idx] = parentV
    mn.data[parent] = value
    mn.heapifyUp(parent)
  }
}

func (mn *MinHeap) heapifyDown(idx int) {
  lidx := mn.getLeftChild(idx)
  ridx := mn.getRightChild(idx)

  if idx >= mn.Length || lidx == mn.Length {
    return
  }

  lval := mn.data[lidx]
  rval := mn.data[ridx]
  val := mn.data[idx]

  if lval > rval && val > rval {
    mn.data[idx] = rval
    mn.data[ridx] = val
    mn.heapifyDown(ridx)
  } else if rval > lval && val > lval {
    mn.data[idx] = lval
    mn.data[lidx] = val
    mn.heapifyDown(ridx)
  }
}

func (mn *MinHeap) parent(idx int) int {
  return int(
    math.Floor(float64(idx - 1) / 2),
  )
}

func (mn *MinHeap) getLeftChild(child int) int {
  return child * 2 + 1
}

func (mn *MinHeap) getRightChild(child int) int {
  return child * 2 + 2
}
