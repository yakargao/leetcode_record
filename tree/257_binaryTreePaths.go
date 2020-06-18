/**
* @Author: CiachoG
* @Date: 2020/5/20 14:36
* @Description：
 */
package algorithm

import "strconv"

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0)
	btPaths(root, "")
	return paths
}
func btPaths(root *TreeNode, pre string) {
	if root == nil {
		return
	}
	if pre != "" {
		pre = pre + "->" + strconv.Itoa(root.Val)
	} else {
		pre = strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		paths = append(paths, pre)
	}
	btPaths(root.Left, pre)
	btPaths(root.Right, pre)
}
