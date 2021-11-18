package sort

/*
注意两个细节优化点：
1. 有序界定，冒泡后边的都是整理好顺序的，不需要再次去遍历（记录最后一次交换的位置即可）
2. 已经排好序的直接退出（只要不在进行交换表示都已排好序）
*/
func Bubble(nums []int) []int {
	var hasSortFlag bool
	sortBorder, lastExchangeIndex := len(nums)-1, len(nums)-1
	for i := 0; i < len(nums)-1; i++ {
		hasSortFlag = true
		for k := 0; k < sortBorder; k++ {
			if nums[k] > nums[k+1] {
				nums[k], nums[k+1] = nums[k+1], nums[k]
				hasSortFlag = false
				lastExchangeIndex = k
			}
		}

		sortBorder = lastExchangeIndex

		if hasSortFlag == true {
			break
		}
	}

	return nums
}
