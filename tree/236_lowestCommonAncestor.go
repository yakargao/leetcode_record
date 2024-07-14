package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//如果都不为空，说明p和q分别在root的左右子树，那么root就是最近公共祖先
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left != nil {
		return left
	}
	return right
}
