package doublePoint

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	res := math.MaxInt
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		right++
		for sum >= right {
			if sum == target {
				res = min(res, right-left)
			}
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		res = 0
	}
	return res

}
