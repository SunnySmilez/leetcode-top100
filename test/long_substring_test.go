package test

import (
	"algo/algo"
	"testing"
)

func TestLongSubString(t *testing.T) {
	type params struct {
		Params string
		Result int
	}

	testP1 := params{
		Params: "abcabcbb",
		Result: 3,
	}

	testP2 := params{
		Params: "bbbbb",
		Result: 1,
	}

	testP3 := params{
		Params: "pwwkew",
		Result: 3,
	}

	testP4 := params{
		Params: "",
		Result: 0,
	}

	testP5 := params{
		Params: "ac",
		Result: 2,
	}

	testP6 := params{
		Params: "c",
		Result: 1,
	}
	testArr := []params{testP1, testP2, testP3, testP4, testP5, testP6}

	for k, v := range testArr {
		res := algo.LongSubString(v.Params)
		if res != v.Result {
			t.Error("Case:", k, "Error, Expected:", v.Result, ",Got:", res)
			break
		}
	}
}
