package levelorder

import "github.com/gogo/protobuf/test/empty-issue70"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	stack := make([]*TreeNode, 0, 1)
	if root == nil {
		return result
	}
	stack = append(stack, root)
	for {
		l := len(stack)
		if l == 0 {
			break
		}
		temp := make([]int, 0, l)
		stack2 := make([]*TreeNode, 0)
		for _, tree := range stack {
			temp = append(temp, tree.Val)
			if tree.Left != nil {
				stack2 = append(stack2, tree.Left)
			}
			if tree.Right != nil {
				stack2 = append(stack2, tree.Right)
			}
		}
		result = append(result, temp)
		stack = stack2
	}

	return result
}
