package cn

import "container/list"

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。 
//
// 返回滑动窗口中的最大值。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], k = 1
//输出：[1]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
// 
//
// 示例 4： 
//
// 
//输入：nums = [9,11], k = 2
//输出：[11]
// 
//
// 示例 5： 
//
// 
//输入：nums = [4,-2], k = 2
//输出：[4] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 105 
// -104 <= nums[i] <= 104 
// 1 <= k <= nums.length 
// 
// Related Topics 堆 Sliding Window 
// 👍 960 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

type MonotonicQueue struct {
	queue list.List
}

func (q *MonotonicQueue)push(val int)  {
	for q.queue.Len() != 0 && q.queue.Back().Value.(int) < val{
		q.queue.Remove(q.queue.Back())
	}
	q.queue.PushBack(val)
}

func (q *MonotonicQueue)pop(n int)  {
	if q.queue.Front().Value.(int) == n { //想删除的元素可能不在单调队列里面了
		q.queue.Remove(q.queue.Front())
	}
}

func (q *MonotonicQueue)max()int {
	return q.queue.Front().Value.(int)
}
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int,0)
	window := &MonotonicQueue{}
	for i:=0;i<len(nums);i++{
		if i < k - 1 {
			window.push(nums[i])
		}else{
			window.push(nums[i])
			result = append(result,window.max())
			window.pop(nums[i-k+1])
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
