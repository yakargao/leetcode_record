package dp

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	//dp[i][j] 用前i个物品,容量为j是否可以刚好装满
	dp := make([][]bool, len(nums)+1)

	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]bool, len(nums)+1)
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= target; j++ {
			if target-nums[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target]

}
