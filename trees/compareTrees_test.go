package trees

import "testing"

func TestCompare(t *testing.T) {
	t1 := BinaryTree{
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

	t2 := BinaryTree{
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

	expected := true
	got := IsEqual(&t1, &t2)

	if expected != got {
		t.Fatalf("Failed compare trees")
	}
}
