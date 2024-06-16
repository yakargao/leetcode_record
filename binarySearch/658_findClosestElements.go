package binarySearch

func findClosestElements(arr []int, k int, x int) []int {
	right := binaryLeftBound2(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
func binaryLeftBound2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
