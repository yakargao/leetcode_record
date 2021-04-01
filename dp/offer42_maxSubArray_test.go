package dp

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	n := 5
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := l + i - 1
			fmt.Printf("[%v,%v]\n", i, j)
		}
	}
}
