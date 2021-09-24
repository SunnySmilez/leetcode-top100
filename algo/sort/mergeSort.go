package sort

/*
	归并排序：将两个有序的数据进行合并成一个有序的数组
*/
func MergeSort(a, b []int) (c []int) {
	c = []int{}

	la := len(a)
	lb := len(b)
	maxL := la + lb
	ak := 0
	bk := 0

	for i := 0; i < maxL; i++ {
		if a[ak] == b[bk] {
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

		if ak == la || bk == lb {
			break
		}
	}

	if ak < len(a) {
		c = append(c, a[ak:]...)
	}

	if bk < len(b) {
		c = append(c, b[bk:]...)
	}

	return
}
