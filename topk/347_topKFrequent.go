package topk

import "math"

type MinHeapV2 struct {
	vals []int
	fm   map[int]int // 出现频率
}

func NewMinHeapV2(fm map[int]int) *MinHeapV2 {
	return &MinHeapV2{
		vals: []int{math.MinInt},
		fm:   fm,
	}
}

// 插入尾部节点然后上浮
func (h *MinHeapV2) Insert(val int) {
	last := len(h.vals)
	h.vals = append(h.vals, val)
	for ; h.fm[h.vals[last/2]] > h.fm[val]; last /= 2 {
		h.vals[last] = h.vals[last/2]
	}
	h.vals[last] = val
}
func (h *MinHeapV2) DeleteMin() int {
	minVal := h.vals[1]
	lastVal := h.vals[len(h.vals)-1]
	child, i := 0, 1
	for ; i*2 < len(h.vals)-1; i = child {
		child = i * 2
		if child < len(h.vals)-1 && h.fm[h.vals[child+1]] < h.fm[h.vals[child]] {
			child++
		}
		if h.fm[lastVal] > h.fm[h.vals[child]] {
			h.vals[i] = h.vals[child]
		} else {
			break
		}
	}
	h.vals[i] = lastVal
	h.vals = h.vals[:len(h.vals)-1]
	return minVal
}
func (h *MinHeapV2) getMin() int {
	return h.vals[1]
}
func (h *MinHeapV2) getSize() int {
	return len(h.vals) - 1
}

func topKFrequent(nums []int, k int) []int {
	frequentMap := make(map[int]int)
	for _, num := range nums {
		frequentMap[num]++
	}
	heap := NewMinHeapV2(frequentMap)
	for num, f := range frequentMap {
		if heap.getSize() < k {
			heap.Insert(num)
		} else {
			if f > frequentMap[heap.getMin()] {
				heap.DeleteMin()
				heap.Insert(num)
			}
		}
	}
	res := make([]int, 0)

	for heap.getSize() != 0 {
		res = append(res, heap.getMin())
		heap.DeleteMin()
	}
	return res
}
