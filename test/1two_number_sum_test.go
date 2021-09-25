package test

import (
	"algo/algo"
	"reflect"
	"testing"
)

type params struct {
	Params1 []int
	Params2 int
	Result  []int
}

func TestTwoNumberSum(t *testing.T) {
	testP1 := params{
		Params1: []int{2, 7, 11, 15},
		Params2: 9,
		Result:  []int{0, 1},
	}

	testP2 := params{
		Params1: []int{3, 2, 4},
		Params2: 6,
		Result:  []int{1, 2},
	}

	testP3 := params{
		Params1: []int{3, 3},
		Params2: 6,
		Result:  []int{0, 1},
	}

	testArr := []params{testP1, testP2, testP3}
	for k, v := range testArr {
		res := algo.TwoNumberSum(v.Params1, v.Params2)
		if reflect.DeepEqual(res, v.Result) == false {
			t.Error("Case:", k, "Error, Expected:", v.Result, ",Got:", res)
		}
	}
}
