package algo

/*
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。



示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
示例 4：

输入：x = -101
输出：false


提示：

-231 <= x <= 231 - 1

*/
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) { // 负数肯定不行，最后一位为0的整数也不行：100，001肯定对不上，因为高位肯定不能是0
		return false
	}

	var r int
	for x > r {
		r = r*10 + x%10
		x = x / 10
	}

	if x == r || r/10 == x { // 偶数的时候直接相等，奇数的时候需要多的除10
		return true
	}

	return false
}
