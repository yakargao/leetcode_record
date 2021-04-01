package dp

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	pre, res := nums[0], nums[0]
	curr := 0
	for i := 1; i < length; i++ {
		num := nums[i]
		if num+pre >= num {
			curr = num + pre
		} else {
			curr = num
		}
		pre = curr
		if curr > res {
			res = curr
		}
	}
	return res
}

func maxSubArray2(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i-1]+nums[i])
		res = max(res, nums[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
