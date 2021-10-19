package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type params struct {
		Params []int
		Result int
	}

	testP1 := params{
		Params: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		Result: 6,
	}

	testP2 := params{
		Params: []int{1},
		Result: 1,
	}

	testP3 := params{
		Params: []int{0},
		Result: 0,
	}

	testP4 := params{
		Params: []int{-1},
		Result: -1,
	}

	testP5 := params{
		Params: []int{-100000},
		Result: -100000,
	}

	testP6 := params{
		Params: []int{-2, 1},
		Result: 1,
	}

	testP7 := params{
		Params: []int{-2, -1},
		Result: -1,
	}

	testP8 := params{
		Params: []int{-1, -2},
		Result: -1,
	}

	testArr := []params{testP1, testP2, testP3, testP4, testP5, testP6, testP7, testP8}

	for k, v := range testArr {
		res := hot100.MaxSubArray(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %+v,Expected:%d,Got:%d", k, v.Params, v.Result, res)
			break
		}
	}
}
