/**
* @Author: CiachoG
* @Date: 2020/6/20 15:00
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

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = math.MinInt32
	dfs(root)
	return ans - 1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, dfs(root.Left))
	right := max(0, dfs(root.Right))
	s1 := 1 + max(0, left) + max(0, right) //left+right+root.Val将左右路径连成一条路线
	s2 := 1 + max(0, max(left, right))     //未root提供最长路线
	ans = max(ans, max(s1, s2))
	return max(left, right) + 1 //返回该节点最大的路径长度
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
