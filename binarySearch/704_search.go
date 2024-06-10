package binarySearch

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}

	}
	return -1
}
