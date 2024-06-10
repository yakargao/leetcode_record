package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	longestUnivaluePathDepth(root)
	return res
}
func longestUnivaluePathDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := longestUnivaluePathDepth(root.Left)
	r := longestUnivaluePathDepth(root.Right)
	ll, rr := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		ll = l + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		rr = r + 1
	}
	res = max(res, ll+rr)
	return max(ll, rr)
}
