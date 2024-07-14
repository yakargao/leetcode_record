package tree

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			if j == len(q)-1 {
				ret = append(ret, node.Val)
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
