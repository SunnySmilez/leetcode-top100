package dynamicProgramming

/*
509. 斐波那契数 https://leetcode-cn.com/problems/fibonacci-number
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给你 n ，请计算 F(n) 。

示例 1：
输入：2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1

示例 2：
输入：3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2

示例 3：
输入：4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3

提示：
0 <= n <= 30
*/

// 递归暴力计算
func Fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// 使用备忘录（map）解决重叠子问题计算
func FibMemo(n int) int {
	if n == 0 {
		return 0
	}

	memoN := map[int]int{}
	return memo(n, memoN)
}

func memo(n int, memoN map[int]int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if v, ok := memoN[n]; v != 0 && ok == true {
		return memoN[n]
	}

	memoN[n] = memo(n-1, memoN) + memo(n-2, memoN)

	return memoN[n]
}

func FibDp(n int) int {
	dp := map[int]int{
		1: 1,
		2: 1,
	}

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
