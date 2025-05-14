package trees

func DfsBst(bst *BinaryTree, n int) bool {
	if bst == nil {
		return false
	} else if bst.Value == n {
		return true
	} else if n < bst.Value {
		return DfsBst(bst.Left, n)
	}
	return DfsBst(bst.Right, n)
}
