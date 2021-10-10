package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type params struct {
		Params int
		Result bool
	}

	testP1 := params{
		Params: 121,
		Result: true,
	}

	testP2 := params{
		Params: -121,
		Result: false,
	}

	testP3 := params{
		Params: 10,
		Result: false,
	}

	testP4 := params{
		Params: -101,
		Result: false,
	}

	testP5 := params{
		Params: 123,
		Result: false,
	}
	testArr := []params{testP1, testP2, testP3, testP4, testP5}

	for k, v := range testArr {
		res := hot100.IsPalindrome(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d,Expected:%t,Got:%t", k, v.Params, v.Result, res)
		}
	}
}
