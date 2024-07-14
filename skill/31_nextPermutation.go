package skill

func nextPermutation(nums []int) {

	swapIdx := len(nums) - 2

	for swapIdx >= 0 && nums[swapIdx] > nums[swapIdx+1] {
		swapIdx--
	}
	if swapIdx >= 0 {
		j := len(nums) - 1

		for j >= swapIdx && nums[j] < nums[swapIdx] {
			j--
		}
		nums[swapIdx], nums[j] = nums[j], nums[swapIdx]
	}
	// 反转 a[i+1:]
	left, right := swapIdx+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
