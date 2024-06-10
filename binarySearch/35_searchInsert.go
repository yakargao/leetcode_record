package binarySearch

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			ans = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}

	}
	return ans
}
