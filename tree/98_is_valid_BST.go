package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	var pre *int
	var res = true
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || !res {
			return
		}
		inorder(root.Left)
		if pre != nil && root.Val <= *pre {
			res = false
		}
		pre = &root.Val
		inorder(root.Right)

	}
	inorder(root)
	return res
}
