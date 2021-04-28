package sort

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 5, 6, 8, 3, 9}
	nums = mergeSort(nums)
	fmt.Println(nums)
}
