package trees

func IsEqual(t1 *BinaryTree, t2 *BinaryTree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Value != t2.Value {
		return false
	}
	return IsEqual(t1.Left, t2.Left) && IsEqual(t1.Right, t2.Right)
}
