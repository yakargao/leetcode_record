package tree

func findTargetNode(root *TreeNode, cnt int) int {
	var currCnt = 0
	var target = 0
	var inorder func(root *TreeNode, cnt int)
	inorder = func(root *TreeNode, cnt int) {
		if currCnt > cnt || root == nil {
			return
		}
		inorder(root.Left, cnt)
		currCnt++
		if currCnt == cnt {
			target = root.Val
		}
		inorder(root.Right, cnt)
	}
	inorder(root, cnt)
	return target
}
