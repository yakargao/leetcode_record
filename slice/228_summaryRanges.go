package slice

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := make([]string, 0)
	start := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if start == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
		}
	}

	// 处理最后一个区间
	if start == nums[len(nums)-1] {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
	}

	return result
}
