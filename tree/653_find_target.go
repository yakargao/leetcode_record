package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	//前序遍历，使用map
	valSet := make(map[int]struct{})
	var preOrder func(root *TreeNode) bool
	preOrder = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := valSet[k-root.Val]; ok {
			return true
		}
		valSet[root.Val] = struct{}{}
		return preOrder(root.Left) || preOrder(root.Right)
	}
	return preOrder(root)
}
