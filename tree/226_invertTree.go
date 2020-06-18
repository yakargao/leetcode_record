/**
* @Author: CiachoG
* @Date: 2020/5/15 16:44
* @Description：
 */
package algorithm

// url :https://leetcode-cn.com/problems/invert-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//自底向上翻转，而不是顶向下翻转
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
