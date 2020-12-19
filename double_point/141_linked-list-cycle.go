package double_point

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	p := head
	m := make(map[*ListNode]struct{})
	for p != nil {
		if _, ok := m[p]; ok {
			return true
		}
		m[p] = struct{}{}

		p = p.Next
	}
	return false
}
