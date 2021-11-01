package hot100

import "sort"

/*
169. 多数元素 : https://leetcode-cn.com/problems/majority-element/

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：
输入：[3,2,3]
输出：3

示例 2：
输入：[2,2,1,1,1,2,2]
输出：2
*/
func MajorityElement(nums []int) int {
	h := map[int]int{}
	for i := 0; i < len(nums); i++ {
		h[nums[i]] += 1
		if h[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}

	return 0
}

func MajorityElement1(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return nums[len(nums)/2]
}
