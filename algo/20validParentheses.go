package algo

import "strings"

/*
20. 有效的括号:https://leetcode-cn.com/problems/valid-parentheses/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false

示例 5：
输入：s = "{[]}"
输出：true

提示：
1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/
func IsValid(s string) bool {
	mapS := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	left := ""
	right := ""
	for _, v := range s {
		if strings.Contains("({[", string(v)) {
			if right != "" {
				if len(left) < len(right) {
					return false
				}

				if left[len(left)-len(right):] == right {
					left = left[0 : len(left)-len(right)]
					right = ""
				} else {
					return false
				}
			}

			left += string(v)
		}

		if v1, ok := mapS[string(v)]; ok {
			if left == "" {
				return false
			}

			right = v1 + right
		}
	}

	if left == right {
		return true
	}

	return false
}
