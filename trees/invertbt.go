package trees

func Invert(bt *BinaryTree) {
	if bt == nil {
		return
	}
	tmp := bt.Left
	bt.Left = bt.Right
	bt.Right = tmp

	Invert(bt.Left)
	Invert(bt.Right)
}
