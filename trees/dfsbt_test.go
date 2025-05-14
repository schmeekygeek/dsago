package trees

import (
	"reflect"
	"testing"
)

func TestPreOrder(t *testing.T) {
	// node
	node := CreateBT()

	// expected
	expected := []int{7, 23, 5, 4, 3, 18, 21}

	// got
	arr := new([]int)
	got := node.DfsPreOrder(arr)
	if !reflect.DeepEqual(expected, got) {
		t.Fatal("PreOrder failed")
	}
}

func CreateBT() BinaryTree {
	node := BinaryTree{
		Left: &BinaryTree{
			Left:  &BinaryTree{Left: nil, Right: nil, Value: 5},
			Right: &BinaryTree{Left: nil, Right: nil, Value: 4},
			Value: 23,
		},
		Right: &BinaryTree{
			Left:  &BinaryTree{Left: nil, Right: nil, Value: 18},
			Right: &BinaryTree{Left: nil, Right: nil, Value: 21},
			Value: 3,
		},
		Value: 7,
	}
	return node
}
