//golang code
package algorithm
 type TreeNode struct {
     Val int
    Left *TreeNode
     Right *TreeNode
 }
func buildTree(preorder []int, inorder []int) *TreeNode {

	//数组中没有节点，直接返回nil
	if len(preorder) == 0  { //len(inorder)==0
		return nil
	}
	//数组中只有一个节点，返回左右子树为nil的节点
	if len(inorder) == 1  {//  len(inorder) == 1
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}
	//数组不止一个节点，递归构造
	idx := indexOfNum(preorder[0],inorder)
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+idx],inorder[:idx]),
		Right: buildTree(preorder[1+idx:],inorder[idx+1:]),
	}
}
func indexOfNum(num int,s []int)int  {
	for idx,val := range s {
		if val == num {
			return idx
		}
	}
	return 0
}
