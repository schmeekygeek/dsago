package trees

type BinaryTree struct {
  Value   int
  Left    *BinaryTree
  Right   *BinaryTree
}

func (root *BinaryTree) DfsPreOrder(path *[]int) []int {
  if root == nil {
    return *path
  }
  *path = append(*path, root.Value)
  root.Left.DfsPreOrder(path)
  root.Right.DfsPreOrder(path)
  return *path
}

func (root *BinaryTree) DfsInOrder(path *[]int) []int {
  if root == nil {
    return *path
  }
  root.Left.DfsInOrder(path)
  *path = append(*path, root.Value)
  root.Right.DfsInOrder(path)
  return *path
}

func (root *BinaryTree) DfsPostOrder(path *[]int) []int {
  if root == nil {
    return *path
  }
  root.Left.DfsPostOrder(path)
  root.Right.DfsPostOrder(path)
  *path = append(*path, root.Value)
  return *path
}

