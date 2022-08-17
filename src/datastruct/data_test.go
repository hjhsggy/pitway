package datastruct

import (
	"testing"
)

func TestBTree(t *testing.T) {

	root := &BinaryTree{}

	root.Root = &TreeNode{Value: 1}
	root.Root.Left = &TreeNode{Value: 2}
	root.Root.Right = &TreeNode{Value: 3}
	root.Root.Right.Left = &TreeNode{Value: 4}
	root.Root.Right.Right = &TreeNode{Value: 5}

	t.Log("PreOrderTraversal\n")
	root.Root.PreOrderTraversal()
	t.Log("MidOrderTraverse\n")
	root.Root.MidOrderTraverse()
	t.Log("BackOrderTraverse\n")
	root.Root.BackOrderTraverse()
	t.Log("LevelTraverse1\n")
	root.Root.LevelTraverse()
	t.Log("LevelTraverse2\n")
	root.Root.LevelTraverse2()

}

/*
         4
   3          5
1     4           7
*/

func TestBSTree(t *testing.T) {

	var root BinaryTree

	root.BSInsert(4)
	root.BSInsert(3)
	root.BSInsert(5)
	root.BSInsert(1)
	root.BSInsert(4)
	root.BSInsert(7)

	t.Log(root.Root.Search(3))

}
