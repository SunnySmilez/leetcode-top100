package hot100

//3. 无重复字符的最长子串:https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// 示例 4:
//
//输入: s = ""
//输出: 0
//
// 提示：
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口

func LongSubString1(s string) (res int) { // 这种思路最大的问题是我每次都需要重新再去从第i个字符开始遍历，实际上中间存在一个已经遍历过的段，将第一个重复字符删除即可继续往后走：比如说skse，如果遍历到第二个s位置，我需要以k为起始位置，继续第二个指针往后遍历也就是kse；而如果我是删除第一个s，第二个指针从e位置开始遍历即可，详细看第二种实现方式
	var hashMap map[byte]int //这种申明方式可以避免""和" "的问题，不要使用map[string]int

	for i := 0; i < len(s); i++ {
		hashMap = map[byte]int{
			s[i]: i,
		}

		for k := i + 1; k < len(s); k++ {
			if _, ok := hashMap[s[k]]; !ok { // 如果是判断存在则获取长度的方式，则需要考虑如果最后也不存在相等的字符且最后的串是最长的这种情况（acb）
				hashMap[s[k]] = k
			}
		}

		if len(hashMap) > res {
			res = len(hashMap)
		}
	}

	return
}

func LongSubString(s string) (res int) {
	hashMap := map[byte]int{}
	rk := 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(hashMap, s[i-1]) //删除前一个，也就是重复的那个
		}

		//_,ok:=hashMap[s[rk]] 这种方式如果写在外边最后一个值会溢出，写在里面会是双层嵌套
		for rk < len(s) && hashMap[s[rk]] == 0 { // 此处直接使用rk来做，且需要注意hashmap不需要使用是否存在的方式来做
			hashMap[s[rk]]++
			rk++
		}

		if len(hashMap) > res {
			res = len(hashMap)
		}
	}

	return res
}
