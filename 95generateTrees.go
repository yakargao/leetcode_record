package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func switchTowNode(node *TreeNode, count int, x int, y int) {
	if node == nil {
		return
	}
	if node.Val == x || node.Val == y {
		if node.Val == x {
			node.Val = y
		} else {
			node.Val = x
		}
		count -= 1
		if count == 0 {
			return
		}
	}
	switchTowNode(node.Left, count, x, y)
	switchTowNode(node.Right, count, x, y)
}
