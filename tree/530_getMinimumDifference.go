/**
* @Author: CiachoG
* @Date: 2020/6/19 21:20
* @Description：
 */
package algorithm

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//二叉搜索树中序遍历的有序性
var diff = 0
var pre = 0

func getMinimumDifference(root *TreeNode) int {
	diff = math.MaxInt32
	pre = math.MinInt32
	getDiff(root)
	return diff
}

func getDiff(root *TreeNode) {
	if root == nil {
		return
	}
	getDiff(root.Left)
	d := abs(root.Val, pre)
	if d < diff {
		diff = d
	}
	pre = root.Val
	getDiff(root.Right)
}
func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
