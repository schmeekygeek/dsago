package trees

func searchBT(bt *BinaryTree, n int) bool {
	if bt == nil {
		return false
	}
	if bt.Value == n {
		return true
	}
	if searchBT(bt.Left, n) {
		return true
	}
	return searchBT(bt.Right, n)
}
