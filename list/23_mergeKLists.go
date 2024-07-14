package list

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

// 二分分治法
func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := l + (r-1)/2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	var head = result
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			result.Next = &ListNode{Val: l2.Val}
			result = result.Next
			l2 = l2.Next
		} else {
			result.Next = &ListNode{Val: l1.Val}
			result = result.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		result.Next = l1
	} else if l2 != nil {
		result.Next = l2
	}
	return head.Next
}
