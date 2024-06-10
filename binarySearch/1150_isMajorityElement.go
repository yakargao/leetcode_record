package binarySearch

func isMajorityElement(nums []int, target int) bool {
	l, r := binaryLeftBound(nums, target), binaryRightBound(nums, target)
	return l != -1 && r != -1 && r-l+1 > len(nums)/2
}

func binaryLeftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l <= len(nums)-1 && nums[l] == target {
		return l
	}
	return -1
}

func binaryRightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if r >= 0 && nums[r] == target {
		return r
	}
	return -1
}
