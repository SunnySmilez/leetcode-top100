package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestReverse(t *testing.T) {
	type params struct {
		Params1 int
		Result  int
	}

	testP1 := params{
		Params1: 123,
		Result:  321,
	}

	testP2 := params{
		Params1: -123,
		Result:  -321,
	}

	testP3 := params{
		Params1: 120,
		Result:  21,
	}

	testP4 := params{
		Params1: 0,
		Result:  0,
	}

	testP5 := params{
		Params1: 2147483648,
		Result:  0,
	}

	testP6 := params{
		Params1: -2147483644,
		Result:  0,
	}

	testP7 := params{
		Params1: 1463847412,
		Result:  2147483641,
	}

	testArr := []params{testP1, testP2, testP3, testP4, testP5, testP6, testP7}

	for k, v := range testArr {
		res := hot100.Reverse(v.Params1)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d,Expected:%d,Got:%d", k, v.Params1, v.Result, res)
			break
		}
	}
}
