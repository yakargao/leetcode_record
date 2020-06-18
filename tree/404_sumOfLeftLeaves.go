/**
* @Author: CiachoG
* @Date: 2020/5/20 15:00
* @Description：
 */
package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			return root.Left.Val + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
		} else {
			return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
		}
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
