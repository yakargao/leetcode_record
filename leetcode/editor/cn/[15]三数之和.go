package cn

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -105 <= nums[i] <= 105 
// 
// Related Topics 数组 双指针 
// 👍 3257 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	result := make([][]int,0)
	for idx,num := range nums{
		//有序
		if num > 0{
			return result
		}
		//去重
		if idx >1 && num == nums[idx-1] {
			continue
		}
		left,right := idx+1,len(nums) - 1
		for left < right{
			sum := num + nums[left] + nums[right]
			if  sum==0 {
				result = append(result,[]int{num,nums[left],nums[right]})
				left ++
				right --
			}else if sum > 0 {
				right --
			}else{
				left ++
			}
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
