package trees

import (
	"testing"
)

func TestBfs(t *testing.T) {
  tree := Tree{
  	Value: 13,
  	nodes: []*Tree{
      {
      	Value: 12,
      	nodes: []*Tree{
          {Value: 7, nodes: []*Tree{}},
          {Value: 8, nodes: []*Tree{}},
        },
      },
      {Value: 4, nodes: []*Tree{}},
      {
      	Value: 1,
      	nodes: []*Tree{
          {Value: 2, nodes: []*Tree{}},
          {Value: 5, nodes: []*Tree{}},

        },
      },
    },
  }
  tree.PrintTree()
  t.Log(tree.Bfs(2))
}
