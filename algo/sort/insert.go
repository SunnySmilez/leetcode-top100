package sort

/**
* 插入排序
* http://notepad.yehyeh.net/Content/Algorithm/Sort/Insertion/1.php
* 思路：在要排序的一组数中，假设前面的数已经是排好顺序的，现在要把第n个数插到前面的有序数中，使得这n个数也是排好顺序的。如此反复循环，直到全部排好顺序
* @param $data
* @return mixed
 */

func Insert(nums []int) []int {
	num := len(nums)

	for i := 1; i < num; i++ {
		for k := i - 1; k >= 0; k-- {
			if nums[k] < nums[i] {
				tmp := nums[i]
				nums[k] = tmp
				nums[k+1] = nums[k]
			}
		}
	}

	return nums
}
