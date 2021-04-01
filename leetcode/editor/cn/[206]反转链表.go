package cn
//反转一个单链表。
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表 
// 👍 1558 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		 return head
	}
	if head.Next == nil { //只剩下一个节点的时候
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next  = nil
	return last
}
//leetcode submit region end(Prohibit modification and deletion)
