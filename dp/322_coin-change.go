package dp

import "math"

/**
1.确认base case：dp[0]=0
2、确认状态：目标金额amount
3、选择：选择一个硬币会减少目标金额
4、dp数组：dp[amount] = min(currMin,1+dp[amount-i]) //i是当前选择的金额
 */

func coinChange(coins []int, amount int) int {
	return  coinChangeFunc2(coins,amount)
}

//递归+备忘录
func coinChangeFunc1(memo map[int]int ,coins []int, amount int) int  {
	if v,ok :=  memo[amount];ok {
		return v
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res :=math.MaxInt32
	for _,coin := range coins{
		sub := coinChangeFunc1(memo,coins,amount-coin)
		if  sub == -1{
			continue
		}
		res = min(res,1+sub)
	}
	if res !=   math.MaxInt32  {
		memo[amount] = res
		return res
	}
	memo[amount] = -1
	return -1
}

//动态规划

func coinChangeFunc2(coins []int, amount int)int{
	 dp := make([]int,amount+1)

	 dp[0]=0
	for i:=1;i<=amount;i++ {
		dp[i] =amount+1
		for _,coin := range coins{
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i],1+dp[i-coin])
		}
	}
	if dp[amount] == amount + 1 {
		return  -1
	}
	return dp[amount]
}

