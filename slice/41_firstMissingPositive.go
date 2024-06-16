package slice

/*
如果数组中出现小于0或者大于等于n+1的数，则必定确实最小正整数[1,n+1]
*/
func firstMissingPositive(nums []int) int {
	l := len(nums)
	has1 := false
	for _, num := range nums {
		if num == 1 {
			has1 = true
			break
		}
	}
	if !has1 {
		return 1
	}

	for i, num := range nums {
		if num <= 0 || num > l {
			nums[i] = 1
		}
	}
	for _, num := range nums {
		idx := abs(num) - 1
		nums[idx] = -abs(nums[idx])
	}
	for idx, num := range nums {
		if num > 0 {
			return idx + 1
		}
	}
	return l + 1
}
func abs(i int) int {
	if i <= 0 {
		i = -i
	}
	return i
}
