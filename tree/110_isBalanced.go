/**
* @Author: CiachoG
* @Date: 2020/5/15 14:23
* @Description：
 */
package algorithm

//https://leetcode-cn.com/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	sub := maxDepth(root.Left) - maxDepth(root.Right)
	return (sub+(sub>>31))^(sub>>31) < 2 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)

}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lmax := maxDepth(root.Left)
	rmax := maxDepth(root.Right)
	if lmax > rmax {
		return lmax + 1
	} else {
		return rmax + 1
	}
}
