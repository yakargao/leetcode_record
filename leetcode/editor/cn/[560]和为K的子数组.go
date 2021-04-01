package cn

d
//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
// 示例 1 : 
//
// 
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
// 
//
// 说明 : 
//
// 
// 数组的长度为 [1, 20,000]。 
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。 
// 
// Related Topics 数组 哈希表 
// 👍 736 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/*
*	前缀和
 */
//func subarraySum(nums []int, k int) int {
//	ans,preSum,target := 0,0,0
//	sumMap := make(map[int]int)
//	sumMap[0] =1
//	for _,num := range nums {
//		preSum += num
//		//计算前面数组从0……i-1这个区间的前缀和为preSum - k
//		target = preSum - k
//		if count,ok := sumMap[target];ok {
//			ans += count
//		}
//		sumMap[preSum]++
//	}
//	return ans
//}
func subarraySum(nums []int, k int) int {
	sum := make([]int,len(nums)+1)
	for idx,num := range nums{
		sum[idx +1] = sum[idx] + num
	}
	ans := 0
	for i := 1;i<len(sum);i++{
		for j := 0;j<i;j++{
			if sum[i] - sum[j] == k{
				ans ++
			}
		}
	}
	return ans
}


//leetcode submit region end(Prohibit modification and deletion)
