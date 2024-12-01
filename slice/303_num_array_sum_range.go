package slice

import "fmt"

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	fmt.Println(sums)
	return NumArray{sums}
}

func (na *NumArray) SumRange(i, j int) int {
	return na.sums[j+1] - na.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
