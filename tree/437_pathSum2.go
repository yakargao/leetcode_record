/**
* @Author: CiachoG
* @Date: 2020/5/20 15:30
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
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return rootPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func rootPath(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	remain := sum - root.Val
	result := 0
	if remain == 0 {
		result = 1
	}
	return result + rootPath(root.Left, remain) + rootPath(root.Right, remain)
}
