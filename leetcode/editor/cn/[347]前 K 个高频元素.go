package cn

import (
	"container/heap"
)

//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
// 
//
// 示例 1: 
//
// 输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
// 
//
// 示例 2: 
//
// 输入: nums = [1], k = 1
//输出: [1] 
//
// 
//
// 提示： 
//
// 
// 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。 
// 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。 
// 你可以按任意顺序返回答案。 
// 
// Related Topics 堆 哈希表 
// 👍 727 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type Item struct {
	Val int
	Num int
}
type MyHeap []Item
func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i].Num > h[j].Num }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	myHeap := make(MyHeap,0)
	for _,num := range nums{
		numMap[num] ++
	}
	for k,v := range numMap{
		myHeap = append(myHeap,Item{Val: k,Num: v})
	}
	heap.Init(&myHeap)
	result := make([]int,k)
	for i:=0;i<k;i++{
		result[i] = heap.Pop(&myHeap).(Item).Val
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
