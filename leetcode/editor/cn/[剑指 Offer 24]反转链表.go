package cn
//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 
//
// 限制： 
//
// 0 <= 节点个数 <= 5000 
//
// 
//
// 注意：本题与主站 206 题相同：https://leetcode-cn.com/problems/reverse-linked-list/ 
// Related Topics 链表 
// 👍 179 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	if head.Next == nil { //只剩下一个节点的时候
//		return head
//	}
//	last := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next  = nil
//	return last
//}
//leetcode submit region end(Prohibit modification and deletion)
