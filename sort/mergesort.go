package sort

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	p1, p2 := 0, 0
	for p1 < len(left) && p2 < len(right) {
		if left[p1] < right[p2] {
			result = append(result, left[p1])
			p1++
		} else {
			result = append(result, right[p2])
			p2++
		}
	}
	if p1 < len(left) {
		result = append(result, left[p1:]...)
	}
	if p2 < len(right) {
		result = append(result, right[p2:]...)
	}
	return result
}
