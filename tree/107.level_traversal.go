/**
* @Author: CiachoG
* @Date: 2020/5/14 19:48
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

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		levelRes := make([]int, 0) //每一层的值的集合
		for i := 0; i < l; i++ {
			node := queue[i]
			levelRes = append(levelRes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append([][]int{levelRes}, result...) //头插法
		queue = queue[l:]
	}
	return result
}
