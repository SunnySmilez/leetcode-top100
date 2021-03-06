package hot100

/*
338. 比特位计数:https://leetcode-cn.com/problems/counting-bits/
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

示例 1：
输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10

示例 2：
输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

提示：
0 <= n <= 105

进阶：
很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）
注意：n=n&(n-1)可计算转二进制后，所含有的1的个数
*/
func CountBits(x int) []int {
	var numsBits []int
	for i := 0; i <= x; i++ {
		numsBits = append(numsBits, numBits(i))
	}

	return numsBits
}

func numBits(n int) (count int) {
	for ; n > 0; n = n & (n - 1) {
		count++
	}

	return count
}

/*
	注意：y & (y−1)=0可判断y是不是2的整数幂
	x比x-high（最近一次2整数幂的值）多1
	bits[i] = bits[i-high]+1
*/
func CountBits1(x int) []int {
	bits := make([]int, x+1)

	high := 0
	for k := 1; k <= x; k++ {
		if k&(k-1) == 0 {
			high = k
		}

		bits[k] = bits[k-high] + 1
	}

	return bits
}
