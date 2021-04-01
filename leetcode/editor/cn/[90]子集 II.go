package cn

import "sort"

//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: [1,2,2]
//输出:
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//] 
// Related Topics 数组 回溯算法 
// 👍 363 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var res1 [][]int
var used map[int]bool
func subsetsWithDup(nums []int) [][]int {
	res1 = make([][]int,0)
	used = make(map[int]bool)
	sort.Ints(nums)
	backTrack1(nums,[]int{},0)
	return res1
}
func backTrack1(nums ,track []int,start int)  {
	tmp := make([]int,len(track))
	copy(tmp,track)
	res1 = append(res1,tmp)
	for i:= start ;i<len(nums);i++{
		if i >0&&nums[i] == nums[i-1]&&!used[nums[i-1]]{
			continue
		}
		track = append(track,nums[i])
		used[nums[i]] = true
		backTrack1(nums,track,i+1)
		used[nums[i]] = false
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
