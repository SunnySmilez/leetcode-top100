package sort

import (
	"algo/algo/sort"
	"reflect"
	"testing"
)

type TestBubbleParams struct {
	Params []int
	Result []int
}

func TestBubble(t *testing.T) {
	test1 := TestBubbleParams{
		Params: []int{0, 3, 2, 1, 6, 4},
		Result: []int{0, 1, 2, 3, 4, 6},
	}

	allTest := []TestBubbleParams{test1}

	for _, v := range allTest {
		res := sort.Bubble(v.Params)

		if reflect.DeepEqual(res, v.Result) == false {
			t.Errorf("params: %+v Expected: %+v,Got:%+v", v.Params, v.Result, res)
		}
	}
}
