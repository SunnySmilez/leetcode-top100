package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestMedian(t *testing.T) {
	type params struct {
		Params1 []int
		Params2 []int
		Result  float64
	}

	testP1 := params{
		Params1: []int{1, 3},
		Params2: []int{2},
		Result:  2.00000,
	}

	testP2 := params{
		Params1: []int{1, 2},
		Params2: []int{3, 4},
		Result:  2.50000,
	}

	testP3 := params{
		Params1: []int{0, 0},
		Params2: []int{0, 0},
		Result:  0.00000,
	}

	testP4 := params{
		Params1: []int{},
		Params2: []int{1},
		Result:  1.00000,
	}
	testP5 := params{
		Params1: []int{2},
		Params2: []int{},
		Result:  2.00000,
	}

	testP6 := params{
		Params1: []int{100000},
		Params2: []int{100001},
		Result:  100000.5,
	}

	testArr := []params{testP1, testP2, testP3, testP4, testP5, testP6}
	for k, v := range testArr {
		res := hot100.Median(v.Params1, v.Params2)
		if res != v.Result {
			t.Error("Case:", k, "Error, Expected:", v.Result, ",Got:", res)
		}
	}
}
