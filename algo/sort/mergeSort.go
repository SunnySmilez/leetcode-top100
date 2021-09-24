package sort

/*
	归并排序：将两个有序的数据进行合并成一个有序的数组
*/
func MergeSort(a, b []int) (c []int) {
	var ak, bk int

	for ak < len(a) && bk < len(b) {
		if a[ak] == b[bk] { //相等两个值都写入
			c = append(c, a[ak])
			c = append(c, b[bk])
			ak = ak + 1
			bk = bk + 1
		} else if a[ak] > b[bk] {
			c = append(c, b[bk])
			bk = bk + 1
		} else if a[ak] < b[bk] {
			c = append(c, a[ak])
			ak = ak + 1
		}
	}

	// 把最后数组的值排在后面
	if ak < len(a) {
		c = append(c, a[ak:]...)
	}

	if bk < len(b) {
		c = append(c, b[bk:]...)
	}

	return
}
