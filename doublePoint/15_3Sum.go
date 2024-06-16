package doublePoint

import "sort"

func threeSum(nums []int) [][]int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sliceLen := len(nums)
	res := make([][]int, 0)
	for i := 0; i < sliceLen; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		tuples := towSum(nums, i+1, sliceLen-1, 0-nums[i])
		for _, t := range tuples {
			t = append(t, nums[i])
			res = append(res, t)
		}
	}
	return res
}

func towSum(nums []int, start, end int, target int) [][]int {
	res := make([][]int, 0)

	for start < end {
		sum := nums[start] + nums[end]
		left := nums[start]
		right := nums[end]
		if sum < target {
			for start < end && nums[start] == left {
				start++
			}
		} else if sum > target {
			for start < end && nums[end] == right {
				end--
			}
		} else {
			res = append(res, []int{nums[start], nums[end]})
			for start < end && nums[start] == left {
				start++
			}
			for start < end && nums[end] == right {
				end--
			}
		}
	}
	return res
}
