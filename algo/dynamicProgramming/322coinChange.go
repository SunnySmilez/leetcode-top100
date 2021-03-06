package dynamicProgramming

import "math"

/*
322. 零钱兑换 https://leetcode-cn.com/problems/coin-change/

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0

示例 4：
输入：coins = [1], amount = 1
输出：1

示例 5：
输入：coins = [1], amount = 2
输出：2


提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
*/
func CoinChange(coins []int, n int) int {
	dict := map[int]int{}
	return coinChangeMemo(dict, coins, n)
}

func coinChangeMemo(dp map[int]int, coins []int, n int) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	if v, ok := dp[n]; ok == true {
		return v
	}

	res := math.MaxFloat64
	for _, coin := range coins {
		subProblem := coinChangeMemo(dp, coins, n-coin)
		if subProblem == -1 {
			continue
		}

		res = math.Min(res, float64(1+subProblem))
		dp[n] = int(res)
	}

	if res == math.MaxFloat64 {
		return -1
	}

	return int(res)
}

func CoinChangeDp(coins []int, amount int) int {
	dp := map[int]float64{
		0: 0,
	}

	for i := 1; i <= amount; i++ {
		dp[i] = float64(amount + 1)
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			dp[i] = math.Min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] == float64(amount+1) {
		return -1
	}

	return int(dp[amount])
}
