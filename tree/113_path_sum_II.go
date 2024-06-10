package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	remainNum := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		if remainNum == 0 {
			return [][]int{{root.Val}}
		}
		return nil
	}
	result := make([][]int, 0)
	leftR := pathSum(root.Left, remainNum)
	for _, r := range leftR {
		nr := append([]int{root.Val}, r...)
		result = append(result, nr)
	}
	rightR := pathSum(root.Right, remainNum)
	for _, r := range rightR {
		nr := append([]int{root.Val}, r...)
		result = append(result, nr)
	}
	return result
}
