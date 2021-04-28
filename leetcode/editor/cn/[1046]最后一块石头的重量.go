package cn

import "fmt"

//有一堆石头，每块石头的重量都是正整数。
//
// 每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下： 
//
// 
// 如果 x == y，那么两块石头都会被完全粉碎； 
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。 
// 
//
// 最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。 
//
// 
//
// 示例： 
//
// 
//输入：[2,7,4,1,8,1]
//输出：1
//解释：
//先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
//再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
//接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
//最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。 
//
// 
//
// 提示： 
//
// 
// 1 <= stones.length <= 30 
// 1 <= stones[i] <= 1000 
// 
// Related Topics 堆 贪心算法 
// 👍 155 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func lastStoneWeight(stones []int) int {

	for i:=len(stones)/2;i>=0;i--{
		adjustDown(stones,i)
	}
	for i:=0;i<len(stones);i++{

		fmt.Println(pop(stones,len(stones)-1-i))
	}
	return 0
}
func abs(a,b int) int  {
	if a > b  {
		return a-b
	}
	return b - a
}
func pop(stones []int,end int) int {
	if len(stones) == 0 {
		return -1
	}
	v := stones[0]
	stones[0],stones[end] = stones[end],stones[0]
	stones = stones[:end]
	adjustDown(stones,0)
	return v
}
func adjustUp(stones []int,n int)  {
	if n <=0 {
		return
	}
	parent := (n-1)/2
	if parent >= 0 && stones[parent]<stones[n] {
		stones[parent],stones[n] = stones[n],stones[parent]
		adjustUp(stones,parent)
	}

}
func adjustDown(stones []int,n int)  {
	l := len(stones)
	if n >= l {
		return
	}
	left := n*2 + 1
	right := n*2 + 2
	maxIdx := n
	if left < l && stones[left] > stones[maxIdx]{
		maxIdx = left
	}
	if right < l && stones[right] > stones[maxIdx]{
		maxIdx = right
	}
	if maxIdx != n {
		stones[n],stones[maxIdx] = stones[maxIdx],stones[n]
		adjustDown(stones,maxIdx)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
