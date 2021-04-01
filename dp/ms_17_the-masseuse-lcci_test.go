package dp

import (
	"fmt"
	"testing"
)

func massage(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		if i-2 >= 0 {
			for j := 0; j <= i-2; j++ {
				dp[i] = max(dp[i], dp[j]+nums[i])
			}
		} else {
			dp[i] = nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
func TestMessage(t *testing.T) {
	fmt.Println(massage([]int{1, 2, 3, 1}))
}
