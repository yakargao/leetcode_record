package binarySearch

import "math"

// 前缀和+2分搜索
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	minLength := math.MaxInt32
	for i := 0; i < n; i++ {
		toFind := prefixSum[i] + target
		index := lowerBound(prefixSum, toFind)
		if index != n+1 {
			minLength = int(math.Min(float64(minLength), float64(index-i)))
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
