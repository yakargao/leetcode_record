package topk

import "testing"

func TestConstructor(t *testing.T) {
	//[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]
	c := Constructor(3, []int{4, 5, 8, 2})
	c.Add(3)
}
