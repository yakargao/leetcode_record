/**
* @Author: CiachoG
* @Date: 2020/6/20 15:49
* @Description：
 */
package algorithm

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for _, c := range root.Children {
		cd := maxDepth1(c)
		if depth < cd {
			depth = cd
		}
	}
	return depth + 1
}
