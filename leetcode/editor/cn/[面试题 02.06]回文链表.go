//编写一个函数，检查输入的链表是否是回文的。 
//
// 
//
// 示例 1： 
//
// 输入： 1->2
//输出： false 
// 
//
// 示例 2： 
//
// 输入： 1->2->2->1
//输出： true 
// 
//
// 
//
// 进阶： 
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
// Related Topics 链表 
// 👍 54 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var first *ListNode
func isPalindrome(head *ListNode) bool {
first = head
return check(head)
}
func check(head *ListNode)bool  {
if head == nil {
return true
}
subRes := check(head.Next)
if first.Val != head.Val {
return false
}
first = first.Next
return subRes
}
//leetcode submit region end(Prohibit modification and deletion)
