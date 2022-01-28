package hot100

/*
49. 字母异位词分组:https://leetcode-cn.com/problems/group-anagrams/
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]
*/
//思路：使用hash：字符串的字母排序后作为key，如果单词内的字母相同，那么排序后的key一定相等
func GroupAnagrams(strs []string) [][]string {
	return nil
}
