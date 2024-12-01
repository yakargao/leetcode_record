package slice

// 异或交换原理：
// a ^= b
// b ^= a
// a ^= b
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// 使用异或操作交换 nums[slow] 和 nums[fast]
			if slow != fast { // 只有在 slow 和 fast 不相同时才进行交换
				nums[slow] ^= nums[fast]
				nums[fast] ^= nums[slow]
				nums[slow] ^= nums[fast]
			}
			slow++
		}
		fast++
	}
}
