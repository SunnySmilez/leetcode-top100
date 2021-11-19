package sort

import (
	"algo/algo/sort"
	"testing"
)

type TestBinarySearchParams struct {
	Params []int
	Target int
	Result int
}

func TestBinarySearch(t *testing.T) {
	test1 := TestBinarySearchParams{
		Params: []int{1, 2, 3, 4, 5},
		Target: 1,
		Result: 0,
	}

	allTest := []TestBinarySearchParams{test1}

	for _, v := range allTest {
		res := sort.BinarySearch(v.Params, v.Target)

		if res != v.Result {
			t.Errorf("params: %+v Expected: %+v,Got:%+v", v.Params, v.Result, res)
		}
	}
}
