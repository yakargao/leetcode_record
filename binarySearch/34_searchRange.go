package binarySearch

func searchRange(nums []int, target int) []int {
	l := binaryLeftBound(nums, target)
	r := binaryRightBound(nums, target)
	return []int{l, r}
}
