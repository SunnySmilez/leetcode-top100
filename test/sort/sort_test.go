package sort

import (
	"algo/algo/sort"
	"reflect"
	"testing"
)

type TestMergeSortParams struct {
	Params1 []int
	Params2 []int
	Result  []int
}

func TestMergeSort(t *testing.T) {
	var allParams []TestMergeSortParams
	test1 := TestMergeSortParams{
		Params1: []int{0},
		Params2: []int{5},
		Result:  []int{0, 5},
	}

	test2 := TestMergeSortParams{
		Params1: []int{1, 3, 5},
		Params2: []int{4, 5, 6},
		Result:  []int{1, 3, 4, 5, 5, 6},
	}

	allTest := append(append(allParams, test1), test2)

	for _, v := range allTest {
		res := sort.MergeSort(v.Params1, v.Params2)

		if reflect.DeepEqual(res, v.Result) == false {
			t.Error("Expected:", v.Result, ",Got:", res)
		}
	}
}
