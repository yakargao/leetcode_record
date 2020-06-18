package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var  fistNode  *TreeNode = nil
var secondNode  *TreeNode = nil
var preNode *TreeNode
func recoverTree(root *TreeNode)  {
	fistNode = nil
	secondNode = nil
	preNode =  nil
	inOrder1(root)
	fistNode.Val = fistNode.Val^secondNode.Val
	secondNode.Val = fistNode.Val^secondNode.Val
	fistNode.Val = fistNode.Val^secondNode.Val
}
func inOrder1(root *TreeNode)  {
	if root == nil {
		return
	}
	inOrder1(root.Left)
	if preNode!=nil && preNode.Val > root.Val {
		if fistNode == nil {
			fistNode = preNode
		}
		secondNode = root
	}
	preNode = root
	inOrder1(root.Right)
}

