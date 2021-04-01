package dp

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		if dp[i] == nil {
			dp[i] = make([]int, m+1)
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text2[i-1] == text1[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
