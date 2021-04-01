package dp

import (
	"fmt"
	"testing"
)

func maxProfit(prices []int) int {
	res := 0
	l := len(prices)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if prices[j]-prices[i] > res {
				res = prices[j] - prices[i]
			}
		}
	}
	return res
}
func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
