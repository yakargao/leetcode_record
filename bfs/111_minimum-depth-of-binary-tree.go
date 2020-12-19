package bfs

import "container/list"

/**
DFS:4 ms	5.1 MB	Go
BFS:272 ms	18 MB	Go
总结：因为dfs实际上是递归，需要遍历整棵树，bfs相当于层序遍历，每一层齐头并进，可以在不遍历整棵树的前提下找到最短距离
但是dfs是空间复杂度是递归栈的深度，最多是数的高度O（logN），而bfs空间复杂度是最多节点层的层数，对于满二叉树，是n/2个节点，即O(N)
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	depth := 1
	queue.PushBack(root)
	for queue.Len() != 0{
		ql := queue.Len()
		for i:=0;i<ql;i++ {
			curr := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			if  curr.Left == nil && curr.Right == nil{
				return depth
			}
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		depth ++
	}
	return depth
}
