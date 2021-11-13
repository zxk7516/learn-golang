package main

import "fmt"

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(value int) *TreeNode {
	return &TreeNode{
		val:   value,
		Left:  nil,
		Right: nil,
	}
}

func FlipTree(t *TreeNode) {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	FlipTree(t.Left)
	FlipTree(t.Right)
}

func TravelTree(t *TreeNode) {
	if t != nil {
		fmt.Print("[ val=", t.val, "\nleft=\n\t")
		TravelTree(t.Left)
		fmt.Print("\n\tright=")
		TravelTree(t.Right)
		fmt.Println("]")
	}

}

func InvertTree(t *TreeNode) {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	InvertTree(t.Left)
	InvertTree(t.Right)
}

func main() {
	tree := NewTree(1)
	tree.Left = NewTree(2)
	tree.Right = NewTree(3)
	tree.Right.Left = NewTree(4)
	tree.Right.Right = NewTree(5)
	TravelTree(tree)
	FlipTree(tree)
	TravelTree(tree)
}
