package dp

// dp[i]从第i间开始
// max(dp[i+1],nums[i]+dp[i+2])
func rob2(nums []int, start, end int) int {
	dp := make([]int, len(nums)+2)
	for i := end; i >= start; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[start]
}
