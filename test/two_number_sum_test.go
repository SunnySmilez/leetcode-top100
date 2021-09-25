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
	params1 := []int{2, 7, 11, 15}
	params2 := 9
	result := []int{0, 1}

	testP1 := params{
		Params1: params1,
		Params2: params2,
		Result:  result,
	}

	params3 := []int{3, 2, 4}
	params4 := 6
	result2 := []int{1, 2}

	testP2 := params{
		Params1: params3,
		Params2: params4,
		Result:  result2,
	}

	params5 := []int{3, 3}
	params6 := 6
	result3 := []int{0, 1}
	testP3 := params{
		Params1: params5,
		Params2: params6,
		Result:  result3,
	}

	var testArr []params
	testArr = append(testArr, testP1)
	testArr = append(testArr, testP2)
	testArr = append(testArr, testP3)

	for _, v := range testArr {
		res := algo.TwoNumberSum(v.Params1, v.Params2)
		if reflect.DeepEqual(res, v.Result) == false {
			t.Error("Expected:", v.Result, ",Got:", res)
		}
	}
}
