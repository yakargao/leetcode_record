package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSumIII(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSumOfRoot(root, targetSum) + pathSumIII(root.Left, targetSum) + pathSumIII(root.Right, targetSum)
}
func pathSumOfRoot(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	remainSum := targetSum - root.Val
	result := 0
	if remainSum == 0 {
		result = 1
	}
	return result + pathSumOfRoot(root.Left, remainSum) + pathSumOfRoot(root.Right, remainSum)
}
