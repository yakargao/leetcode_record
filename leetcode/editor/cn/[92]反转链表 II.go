package cn
//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明: 
//1 ≤ m ≤ n ≤ 链表长度。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL 
// Related Topics 链表 
// 👍 704 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head,right)
	}
	head.Next = reverseBetween(head.Next,left-1,right-1)
	return head
}
var successor *ListNode
func reverseN(head *ListNode,n int)*ListNode  {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next,n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
//leetcode submit region end(Prohibit modification and deletion)
