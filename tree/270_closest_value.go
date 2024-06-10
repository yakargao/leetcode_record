package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var minDiff = math.MaxFloat64
var closestV = 0

func closestValue(root *TreeNode, target float64) int {
	minDiff = math.MaxFloat64
	closestV = 0
	inorder(root, target)
	return closestV
}
func inorder(root *TreeNode, target float64) {
	if root == nil {
		return
	}
	inorder(root.Left, target)
	diff := math.Abs(float64(root.Val) - target)
	if diff < minDiff {
		minDiff = diff
		closestV = root.Val
	}
	inorder(root.Right, target)
}
