package cn
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 
//
// 示例： 
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
// 
// Related Topics 字符串 回溯算法 
// 👍 1514 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
括号合法性：1：整个括号字符串，左括号个数==右括号的个数
         2：局部括号字符串，左括号个数>=右括号个数
*/
var gpRes []string
func generateParenthesis(n int) []string {
	gpRes = make([]string,0)
	gpBackTrack(n,n,[]rune{})
	return gpRes
}
func gpBackTrack(left,right int,track []rune)  {
	if left<0||right<0 {
		return
	}
	//如果左括号剩余过多
	if right < left {
		return
	}
	if left == right && left == 0 {
		gpRes = append(gpRes,string(track))
	}
	track = append(track,'(')
	gpBackTrack(left-1,right,track)
	track = track[:len(track)-1]

	track = append(track,')')
	gpBackTrack(left,right-1,track)
	track = track[:len(track)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
