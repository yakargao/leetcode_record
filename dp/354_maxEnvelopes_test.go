package dp

import (
	"fmt"
	"sort"
	"testing"
)

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	//按照宽度递增，宽度相同按照高度递减
	//为什么是递减，因为宽度相同不能互套
	sort.SliceStable(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})
	//fmt.Println(envelopes)
	height := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		height[i] = envelopes[i][1]
	}
	return LISofMaxEnvelopes(height)
}

//求height的最长递增序列
//O(n^2)
func LISofMaxEnvelopes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func TestMaxEnvelopes(t *testing.T) {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println(maxEnvelopes(envelopes))
}
