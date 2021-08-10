package datastruct

import "fmt"

func (tree *BinaryTree) BSInsert(data int) {

	if tree.Root == nil {
		tree.Root = &TreeNode{
			Value: data,
			Left:  nil,
			Right: nil,
		}
		return
	}

	tree.Root.BSInsert(data)

}

func (node *TreeNode) BSInsert(data int) {

	// 插到左子树
	if data <= node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: data, Left: nil, Right: nil}
		} else {
			node.Left.BSInsert(data)
		}

	} else { // 右子树
		if node.Right == nil {
			node.Right = &TreeNode{Value: data, Left: nil, Right: nil}
		} else {
			node.Right.BSInsert(data)
		}
	}

}

func (node *TreeNode) Search(data int) int {

	if node == nil {
		return -1
	}

	if data == node.Value {
		return 0
	}
	if data > node.Value {
		return node.Right.Search(data)
	}
	if data < node.Value {
		return node.Left.Search(data)
	}

	return -1
}

func (root *TreeNode) PreOrderTraversal() {
	if root != nil {
		root.Left.PreOrderTraversal()
		fmt.Printf("value: %d ", root.Value)
		root.Right.PreOrderTraversal()
	}
}

func (root *TreeNode) MidOrderTraverse() {
	if root != nil {
		fmt.Printf("value: %d ", root.Value)
		root.Left.MidOrderTraverse()
		root.Right.MidOrderTraverse()
	}
}

func (root *TreeNode) BackOrderTraverse() {

	if root != nil {
		root.Right.BackOrderTraverse()
		fmt.Printf("value: %d ", root.Value)
		root.Left.BackOrderTraverse()
	}
}
