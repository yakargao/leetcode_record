package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head
	for i := 0; i < k; i++ {
		if right == nil {
			return head
		}
		right = right.Next
	}
	newHead := reverseBetween(left, right)
	left.Next = reverseKGroup(right, k)
	return newHead

}

func reverseBetween(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
