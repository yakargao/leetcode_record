package cn
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
// 示例: 
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
// Related Topics 回溯算法 
// 👍 467 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var r [][]int
func combine(n int, k int) [][]int {
	r = make([][]int,0)
	backTrack2(n,k,1,[]int{})
	return r
}
func backTrack2(n,k,start int,track []int)  {
	if len(track) == k {
		tmp := make([]int,len(track))
		copy(tmp,track)
		r = append(res,tmp)
	}
	for i:=start;i<=n;i++{
		track = append(track,i)
		backTrack2(n,k,i+1,track)
		track = track[:len(track)-1]
	}

}
//leetcode submit region end(Prohibit modification and deletion)
