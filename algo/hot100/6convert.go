package hot100

import "strings"

/*
6. Z 字形变换:https://leetcode-cn.com/problems/zigzag-conversion/
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I

示例 3：
输入：s = "A", numRows = 1
输出："A"


提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000
https://www.cnblogs.com/springfor/p/3889414.html
*/
/*

P     I    N
A   L S  I G
Y A   H R
P     I
思路：从上往下遍历的时候，每个值是顺着下来，从下往上遍历的时候，每个值是跳一格上去
注意这里不要用二位数组的思路去解，直接按照一维数组加字符串拼接的方式即可
*/
func Convert(s string, numRows int) (res string) {
	if numRows == 1 {
		return s
	}

	k, flag := 0, 0
	arr := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		arr[k] = arr[k] + string(s[i])
		if flag == 0 {
			k++
		} else {
			k--
		}

		if k == numRows-1 {
			flag = 1
		}

		if k == 0 {
			flag = 0
		}
	}

	return strings.Join(arr, "")
}
