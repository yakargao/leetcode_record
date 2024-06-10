package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var preVal, count, maxCount int
var answer []int

func findMode(root *TreeNode) []int {
	preVal, count, maxCount = 0, 0, 0
	answer = make([]int, 0)
	findInorder(root)
	return answer
}
func findInorder(root *TreeNode) {
	if root == nil {
		return
	}
	findInorder(root.Left)

	if root.Val == preVal {
		count++
	} else {
		preVal, count = root.Val, 1
	}
	if count == maxCount {
		answer = append(answer, root.Val)
	} else if count > maxCount {
		maxCount = count
		answer = []int{root.Val}
	}
	findInorder(root.Right)
}
