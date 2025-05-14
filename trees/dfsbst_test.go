package trees

import "testing"

func TestDfsBst(t *testing.T) {
	node := BinaryTree{
		Left: &BinaryTree{
			Left:  &BinaryTree{Left: nil, Right: nil, Value: 11},
			Right: &BinaryTree{Left: nil, Right: nil, Value: 14},
			Value: 13,
		},
		Right: &BinaryTree{
			Left:  &BinaryTree{Left: nil, Right: nil, Value: 16},
			Right: &BinaryTree{Left: nil, Right: nil, Value: 20},
			Value: 18,
		},
		Value: 15,
	}
	expected := true
	got := DfsBst(&node, 14)

	if expected != got {
		t.Fatal("TestDfsBst failed")
	}
}
