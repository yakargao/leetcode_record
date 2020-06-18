/**
* @Author: CiachoG
* @Date: 2020/5/15 15:09
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lMin := minDepth(root.Left)
	rMin := minDepth(root.Right)
	if lMin == 0 {
		return rMin + 1
	}
	if rMin == 0 {
		return lMin + 1
	}
	if lMin < rMin {
		return lMin + 1
	} else {
		return rMin + 1
	}
}
