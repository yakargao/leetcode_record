

//给你一个整数数组 nums ，返回该数组所有可能的子集（幂集）。解集不能包含重复的子集。
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0]
//输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// 
// Related Topics 位运算 数组 回溯算法 
// 👍 948 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var res [][]int
func subsets(nums []int) [][]int {
	res = make([][]int,0)
	backTrack(nums,[]int{},0)
	return res
}
func backTrack(nums ,track []int,start int)  {
	tmp := make([]int,len(track))
	copy(tmp,track)
	res = append(res,tmp)
	for i:= start ;i<len(nums);i++{
		track = append(track,nums[i])
		backTrack(nums,track,i+1)
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
