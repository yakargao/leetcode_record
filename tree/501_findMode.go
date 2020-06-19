/**
* @Author: CiachoG
* @Date: 2020/6/19 20:43
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

var maxCount = 0
var countMap = make(map[int]int)
var countAttr = make(map[int][]int)

func findMode(root *TreeNode) []int {
	maxCount = 0
	countMap = make(map[int]int)
	countAttr = make(map[int][]int)
	inOrderFindMode(root)
	return countAttr[maxCount]
}
func inOrderFindMode(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderFindMode(root.Left)
	countMap[root.Val] = countMap[root.Val] + 1
	countAttr[countMap[root.Val]] = append(countAttr[countMap[root.Val]], root.Val)
	if countMap[root.Val] > maxCount {
		maxCount = countMap[root.Val]
	}
	inOrderFindMode(root.Right)
}
