/**
* @Author: CiachoG
* @Date: 2020/6/20 15:58
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

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lt := findTilt(root.Left)
	rt := findTilt(root.Right)
	t := abs(subNodeSum(root.Left), subNodeSum(root.Right))
	return t + lt + rt
}
func subNodeSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := subNodeSum(node.Left)
	r := subNodeSum(node.Right)
	return l + r + node.Val
}

//需要优化上面的代码
/*
执行用时：28 ms, 在所有 Go 提交中击败了11.59%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了100.00%的用户
*/

var tilt = 0

func findTilt2(root *TreeNode) int {
	tilt = 0
	traverse(root)
	return tilt
}
func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := traverse(root.Left)
	r := traverse(root.Right)
	tilt += abs(l, r)
	return l + r + root.Val
}
