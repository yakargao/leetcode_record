package cn
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1： 
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
// 
//
// 示例 2： 
//
// 输入: "cbbd"
//输出: "bb"
// 
// Related Topics 字符串 动态规划 
// 👍 3037 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	res := s[0:1]
	l := len(s)
	dp := make([][]int,l)
	for i:=0;i<l;i++{
		if dp[i] == nil {
			dp[i] = make([]int,l)
		}
		dp[i][i] = 1
	}
	max := func(a,b int)int {
		if a>b {
			return  a
		}
		return b
	}
	for i:=l-2;i>=0;i-- {
		for j:=i+1;j<l;j++ {
			if s[i]==s[j] {
				dp[i][j] = dp[i+1][j-1]+2
			}else{
				dp[i][j] = max(dp[i+1][j],dp[i][j-1])
			}
			if dp[i][j]> len(res) {
				res = s[i:j+1]
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
