package sort

/**
 * 快速排序
 * https://www.jianshu.com/p/1bd09e10f6db
 * 思路：选择一个基准元素，通常选择第一个元素或者最后一个元素。
 *      通过一趟扫描，将待排序列分成两部分，一部分比基准元素小，一部分大于等于基准元素。
 *      此时基准元素在其排好序后的正确位置，然后再用同样的方法递归地排序划分的两部分
 * @param $params
 * @return mixed
 */
func Quick(nums []int) (res []int) {
	var left []int
	var right []int
	if len(nums) <= 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			right = append(right, nums[i])
		} else {
			left = append(left, nums[i])
		}
	}

	left = Quick(left)
	right = Quick(right)

	left = append(left, nums[0])

	return MergeSort(left, right)
}
