package cn

import "sort"

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。 
//
// 说明： 
//
// 
// 所有数字（包括目标数）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
// 
//
// 示例 2: 
//
// 输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//] 
// Related Topics 数组 回溯算法 
// 👍 470 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var cbs2 [][]int
var used2 map[int]bool
func combinationSum2(candidates []int, target int) [][]int {
	cbs2 = make([][]int,0)
	sort.Ints(candidates)
	used2 = make(map[int]bool)
	combinationSum2Bt(candidates,[]int{},0,target)
	return cbs2
}
func combinationSum2Bt(candidates,track []int,start, target int)  {
	if sliceSum2(track) >= target {
		if sliceSum2(track) == target {
			tmp := make([]int,len(track))
			copy(tmp,track)
			cbs2 = append(cbs2,tmp)
		}
		return
	}
	for i := start;i<len(candidates);i++ {
		if len(track) >=1 && candidates[i] < track[len(track)-1] {
			continue
		}
		if i>0 && candidates[i-1]==candidates[i] && !used2[candidates[i-1]]{
			continue
		}
		track = append(track,candidates[i])
		used2[candidates[i]] = true
		combinationSum2Bt(candidates,track,i+1,target)
		used2[candidates[i]] = false
		track = track[:len(track)-1]
	}
}
func sliceSum2(track []int)int  {
	sum := 0
	for _,num := range track{
		sum += num
	}
	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
