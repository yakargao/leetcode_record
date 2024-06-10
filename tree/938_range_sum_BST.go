package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	var sum = 0
	var inorder func(root *TreeNode, low int, high int)
	inorder = func(root *TreeNode, low int, high int) {
		if root == nil {
			return
		}
		inorder(root.Left, low, high)
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
		inorder(root.Right, low, high)
	}
	inorder(root, low, high)
	return sum
}
