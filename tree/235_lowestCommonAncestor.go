/**
* @Author: CiachoG
* @Date: 2020/5/20 13:51
* @Description：
 */
package algorithm

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

//### 方法1：递归
//LCA处于p,q中间，也就是pq在lca的左右子树上
//三种情况：
// 1. p,q在左子树上
// 2.p,q在右子树上
// 3.p,q分别在左右子树上
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

//执行用时 : 20 ms
//内存消耗 : 6.8 MB
//时间复杂度：O(N)
//其中 N 为 BST 中节点的个数，在最坏的情况下我们可能需要访问 BST 中所有的节点。
//
//空间复杂度：O(N)
//所需开辟的额外空间主要是递归栈产生的，之所以是 N 是因为 BST 的高度为 N

//### 方法二
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	node := root
	for node != nil {
		pVal := node.Val
		if pVal < p.Val && pVal < q.Val {
			node = node.Right
		} else if pVal > p.Val && pVal > q.Val {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}

//执行用时 :24 ms
//内存消耗 :6.5 MB
