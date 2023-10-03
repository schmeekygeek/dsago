package heap

import "fmt"

type MinHeap struct {
  length    int
  data      []int
}

func (m *MinHeap) Insert(data int) {
  m.data = append(m.data, data)
  m.heapifyUp(m.length)
  m.length++
}

func (m *MinHeap) Delete() int {
  if len(m.data) == 0 {
    return -1
  }

  out := m.data[0]
  m.length--
  if len(m.data) == 0 {
    return out
  }

  m.data[0] = m.data[m.length]
  m.heapifyDown(0)
  m.data = m.data[:m.length]
  return out
}


func (m *MinHeap) heapifyDown(idx int) {
  lidx := m.leftChild(idx)
  ridx := m.rightChild(idx)

  if idx >= len(m.data) || lidx >= len(m.data) {
    return
  }

  lval := m.data[lidx]
  rval := m.data[ridx]
  val := m.data[idx]

  if lval > rval && val > rval {
    m.data[ridx] = val
    m.data[idx] = rval
    m.heapifyDown(ridx)
  } else if rval > lval && val > lval {
    m.data[lidx] = val
    m.data[idx] = lval
    m.heapifyDown(lidx)
  }

}

func (m *MinHeap) heapifyUp(idx int) {
  if idx == 0 {
    return
  }
  p := m.parent(idx)
  pvalue := m.data[p]
  value := m.data[idx]

  if pvalue > value {
    m.data[idx] = pvalue
    m.data[p] = value
    m.heapifyUp(p)
  }
}

func (m *MinHeap) parent(idx int) int {
  return (idx-1)/2
}

func (m *MinHeap) leftChild(idx int) int {
  return (idx * 2) + 1
}

func (m *MinHeap) rightChild(idx int) int {
  return (idx * 2) + 2
}

func (m *MinHeap) printHeap() {
  for _, val := range m.data {
    fmt.Println(val)
  }
}
