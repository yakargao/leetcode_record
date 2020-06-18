/**
* @Author: CiachoG
* @Date: 2020/5/15 13:36
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l := len(nums)
	mid := l >> 1
	node := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[0:mid]),
		Right: nil,
	}
	if mid+1 < l {
		node.Right = sortedArrayToBST(nums[mid+1 : l])
	}
	return node
}
