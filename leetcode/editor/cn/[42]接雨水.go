package cn

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
// 
//
// 示例 2： 
//
// 
//输入：height = [4,2,0,3,2,5]
//输出：9
// 
//
// 
//
// 提示： 
//
// 
// n == height.length 
// 0 <= n <= 3 * 104 
// 0 <= height[i] <= 105 
// 
// Related Topics 栈 数组 双指针 动态规划 
// 👍 1943 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	max := func(a,b int)int {
		if a>b {
			return a
		}
		return b
	}
	min := func(a,b int)int {
		if a>b {
			return b
		}
		return a
	}
	//lmax[i]表示[0..i]最高的柱子
	lMax,rMax := make([]int,l),make([]int,l)
	lMax[0],rMax[l-1] = height[0],height[l-1]
	for i:=1;i<l;i++{
		lMax[i] = max(lMax[i-1],height[i])
	}
	for i:=l-2;i>=0;i--{
		rMax[i] = max(rMax[i+1],height[i])
	}
	ans := 0
	for i:=1;i<l-1;i++ {
		ans += min(lMax[i],rMax[i]) - height[i]
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)