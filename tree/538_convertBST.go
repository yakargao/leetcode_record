/**
* @Author: CiachoG
* @Date: 2020/6/20 14:10
* @Description： 把二叉搜索树转换为累加树
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
var pre2 = 0

func convertBST(root *TreeNode) *TreeNode {
	pre2 = 0
	convert(root)
	return root
}
func convert(root *TreeNode) {
	if root == nil {
		return
	}
	convert(root.Right)
	root.Val = root.Val + pre2
	pre2 = root.Val
	convert(root.Left)
}
