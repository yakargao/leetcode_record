package dp

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T)  {
	coins := []int{2}
	fmt.Println(coinChange(coins,3))
}
