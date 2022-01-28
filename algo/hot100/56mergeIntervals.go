package hot100

/*
56. 合并区间:https://leetcode-cn.com/problems/merge-intervals/

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
//解题思路：先进行排序。再根据小区间的右侧元素（left）和大区间的左侧元素（right）进行比对，如果left>right则可合并，否则不可合并，则将比对的小区间替换成当前的大区间的值，与下一个区间进行比对
func Merge(intervals [][]int) [][]int {
	return nil
}
