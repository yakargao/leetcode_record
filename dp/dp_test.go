package dp

import "testing"

func Test(t *testing.T) {
	nums := []int{1, 2}
	res := 0
	if len(nums) == 1 {
		res = nums[0]
	} else {
		res = max(rob2(nums, 1, len(nums)-1), rob2(nums, 0, len(nums)-2))
	}
	t.Log(res)
}
