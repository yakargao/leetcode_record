package cn

import (
	"container/list"
)

//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
// 
//
// 示例： 
//
// 输入：[1,2,3,4,5,null,7,8]
//
//        1
//       /  \ 
//      2    3
//     / \    \ 
//    4   5    7
//   /
//  8
//
//输出：[[1],[2,3],[4,5,7],[8]]
// 
// Related Topics 树 广度优先搜索 
// 👍 44 👎 0
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
type ListNode struct {
	Val int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

func listOfDepth(tree *TreeNode) []*ListNode {
	ln := make([]*ListNode,0)
	queue := list.New()
	queue.PushBack(tree)
	for queue.Len() != 0{
		ql := queue.Len()
		front := &ListNode{}
		t := front
		for i:=0;i<ql;i++{//每一层
			curr,ok  := queue.Remove(queue.Front()).(*TreeNode)
			if !ok {
				panic("assert error")
			}
			t.Next = &ListNode{Val: curr.Val}
			t = t.Next
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		ln = append(ln,front.Next)
	}

	return ln
}
//leetcode submit region end(Prohibit modification and deletion)
