package trees

import "fmt"

type Tree struct {
  Value int
  nodes []*Tree
}

func (root *Tree) Bfs(n int) bool {
  queue := make([]*Tree, 0)
  queue = append(queue, root)

  for len(queue) != 0 {
    node := queue[0]
    queue = queue[1:]
    if node.Value == n {
      return true
    }
    for _, child := range node.nodes {
      queue = append(queue, child)
    }
  }

  return false
}

func (root *Tree) PrintTree() {
  queue := make([]Tree, 0)
  queue = append(queue, *root)
  for len(queue) != 0 {
    node := queue[0]
    queue = queue[1:]
    fmt.Print(node.Value)
    for _, val := range node.nodes {
      queue = append(queue, *val)
    }
    fmt.Println()
  }
}
