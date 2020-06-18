package algorithm

import "math"

//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
// }

var ans = math.MinInt32
func maxPathSum(root *TreeNode) int {
	dfs(root)
	return ans
}

func dfs(root *TreeNode)int{
	if root == nil {
		return 0
	}
	left := max(0,dfs(root.Left))
	right :=max(0,dfs(root.Right))
	s1 := root.Val + max(0,left) + max(0,right) //left+right+root.Val将左右路径连成一条路线
	s2 := root.Val+max(0,max(left,right))//未root提供最长路线
	ans = max(ans,max(s1,s2))
	return max(left,right)+root.Val//返回该节点最大的路径长度
}
func max(a, b int)int{
	if a>b {
		return a
	}
	return b
}
