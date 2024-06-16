package dp

func longestPalindromeSubseq(s string) int {
	//dp[i][j] 表示下标以i开始以j结束的字符串最长的回文子序列
	// i=j 则最长回文子序列是1
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	//dp[i][j] = {
	//	if s[i] == s[j] dp[i+1]dp[j-1] + 1
	//	else max()
	//}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 //注意这里是+2 全部包含
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]

}
