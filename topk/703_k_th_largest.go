package topk

import "math"

type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{arr: []int{math.MinInt}}
}

func (h *MinHeap) Insert(val int) {
	var i = len(h.arr)
	h.arr = append(h.arr, val)
	for ; h.arr[i/2] > val; i = i / 2 {
		h.arr[i] = h.arr[i/2]
	}
	h.arr[i] = val
}

func (h *MinHeap) DeleteMin() int {
	minVal := h.arr[1]
	last := h.arr[len(h.arr)-1]
	child, i := 0, 1
	for ; i*2 <= len(h.arr)-1; i = child {
		child = i * 2
		if child < len(h.arr)-1 && h.arr[child+1] < h.arr[child] {
			child++
		}
		if last > h.arr[child] {
			h.arr[i] = h.arr[child]
		} else {
			break
		}
	}
	h.arr[i] = last
	h.arr = h.arr[:len(h.arr)-1]
	return minVal
}
func (h *MinHeap) getMin() int {
	return h.arr[1]
}
func (h *MinHeap) getSize() int {
	return len(h.arr) - 1
}

type KthLargest struct {
	k    int
	heap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := NewMinHeap()
	for _, num := range nums {
		h.Insert(num)
		if h.getSize() > k {
			h.DeleteMin()
		}
	}
	return KthLargest{k: k, heap: h}
}

func (this *KthLargest) Add(val int) int {
	this.heap.Insert(val)
	if this.heap.getSize() > this.k {
		this.heap.DeleteMin()
	}
	return this.heap.getMin()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
