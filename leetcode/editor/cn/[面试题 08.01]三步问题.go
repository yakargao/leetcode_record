package cn
//三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模100
//0000007。 
//
// 示例1: 
//
// 
// 输入：n = 3 
// 输出：4
// 说明: 有四种走法
// 
//
// 示例2: 
//
// 
// 输入：n = 5
// 输出：13
// 
//
// 提示: 
//
// 
// n范围在[1, 1000000]之间 
// 
// Related Topics 动态规划 
// 👍 35 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func waysToStep(n int) int {
	memo := make([]int,n+1)
	return ways(memo,n)
}
func ways(memo []int,n int)int  {
	if memo[n] != 0 {
		return memo[n]
	}
	switch n {
	case 1:
		memo[n] =  1
	case 2:
		memo[n] =  2
	case 3:
		memo[n] =  4
	default:
		memo[n] =  (ways(memo,n-1)+(ways(memo,n-2)+ways(memo,n-3))% 1000000007)% 1000000007

	}
	return memo[n]
}
//leetcode submit region end(Prohibit modification and deletion)
