package dp

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal, preMaxVal, preMinVal := math.MinInt, 1, 1

	for _, num := range nums {
		if num < 0 {
			preMaxVal, preMinVal = preMinVal, preMaxVal
		}
		preMaxVal = max(num, num*preMaxVal)
		preMinVal = min(num, num*preMinVal)
		maxVal = max(maxVal, preMaxVal)
	}

	return maxVal
}
