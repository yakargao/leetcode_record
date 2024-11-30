package slice

// 摩尔投票法
func majorityElement(nums []int) int {
	val, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			cnt++
			val = num
		} else {
			if val == num {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return val
}
