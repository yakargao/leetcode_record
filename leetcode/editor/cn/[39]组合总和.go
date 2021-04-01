package cn

import "sort"

//给定一个e无重复元素的数组 candidats 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。 
//
// 说明： 
//
// 
// 所有数字（包括 target）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1： 
//
// 输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
// 
//
// 示例 2： 
//
// 输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//] 
//
// 
//
// 提示： 
//
// 
// 1 <= candidates.length <= 30 
// 1 <= candidates[i] <= 200 
// candidate 中的每个元素都是独一无二的。 
// 1 <= target <= 500 
// 
// Related Topics 数组 回溯算法 
// 👍 1115 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var sumRes [][]int
func combinationSum(candidates []int, target int) [][]int {
	sumRes = make([][]int,0)
	sort.Ints(candidates)
	combinationSumBt(candidates,[]int{},target)
	return sumRes
}
func combinationSumBt(candidates,track []int,target int)  {
	if  sliceSum(track) >= target{
		if sliceSum(track) == target {
			tmp := make([]int,len(track))
			copy(tmp,track)
			sumRes = append(sumRes,tmp)
		}
		return
	}
	for _,c := range candidates {
		if len(track)>=1&&c<track[len(track)-1] {
			continue
		}
		track = append(track,c)
		combinationSumBt(candidates,track,target)
		track = track[:len(track)-1]
	}
}
func sliceSum(track []int)int  {
	sum := 0
	for _,num := range track{
		sum += num
	}
	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
