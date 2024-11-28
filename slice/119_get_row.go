package slice

func getRow(rowIndex int) []int {
	dp := make([]int, rowIndex+1)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i <= rowIndex; i++ {
		for j := i - 1; j > 0; j-- {
			dp[j] += dp[j-1]
		}
	}
	return dp
}
