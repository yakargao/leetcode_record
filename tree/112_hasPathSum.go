/**
* @Author: CiachoG
* @Date: 2020/5/15 16:28
* @Description：
 */
package algorithm

//https://leetcode-cn.com/problems/path-sum/submissions/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
 极端情况输入: [] 0
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum = sum - root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
