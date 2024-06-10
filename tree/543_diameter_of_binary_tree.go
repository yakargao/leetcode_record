package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxNodes int

func diameterOfBinaryTree(root *TreeNode) int {
	maxNodes = 1
	diameterOfDepth(root)
	return maxNodes - 1
}
func diameterOfDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := diameterOfDepth(root.Left)
	r := diameterOfDepth(root.Right)
	maxNodes = max(maxNodes, l+r+1)
	return max(l, r) + 1
}
