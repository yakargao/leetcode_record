package dp

// dp[i]从第i间开始
// max(dp[i+1],nums[i]+dp[i+2])
func rob(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}
