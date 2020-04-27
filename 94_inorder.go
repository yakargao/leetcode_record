package algorithm
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var  orderArray []int
func inorderTraversal(root *TreeNode) []int {
	orderArray = make([]int,0)
	inOrder(root)
	return orderArray
}
func inOrder(root *TreeNode)  {
	if root == nil {
		return
	}
	inOrder(root.Left)
	orderArray = append(orderArray,root.Val)
	inOrder(root.Right)
}
