package sort

func quickSort(nums []int, low, high int) {
	if low < high {
		pivotPos := partition(nums, low, high)
		//每一轮有一个数确定了位置
		quickSort(nums, low, pivotPos-1)
		quickSort(nums, pivotPos+1, high)
	}
}
func partition(nums []int, low, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= pivot {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}
