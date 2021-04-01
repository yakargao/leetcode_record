package cn

import "container/list"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串 
// 👍 2096 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	leftStack := list.New()
	for _,ch := range  s {
		if ch=='(' ||ch=='[' || ch=='{' {
			leftStack.PushFront(ch)
		}else{
			if leftStack.Len() != 0 {
				top := leftStack.Front().Value.(rune)
				if top == leftOf(ch)  {
					leftStack.Remove(leftStack.Front())
				}else{
					return false
				}
			}else{
				return false
			}
		}
	}
	return leftStack.Len() == 0
}
func leftOf(ch rune) rune {
	switch ch {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	}
	return ' '
}
//leetcode submit region end(Prohibit modification and deletion)
