package dynamicProgramming

import (
	"algo/algo/dynamicProgramming"
	"testing"
)

type params509 struct {
	Params int
	Result int
}

func TestFib(t *testing.T) {
	testArr := getTestArr509()
	for k, v := range testArr {
		res := dynamicProgramming.Fib(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d,Expected:%d,Got:%d", k, v.Params, v.Result, res)
		}
	}
}

func TestFibMemo(t *testing.T) {
	testArr := getTestArr509()
	for k, v := range testArr {
		res := dynamicProgramming.FibMemo(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d,Expected:%d,Got:%d", k, v.Params, v.Result, res)
		}
	}
}

func TestFibDp(t *testing.T) {
	testArr := getTestArr509()
	for k, v := range testArr {
		res := dynamicProgramming.FibDp(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d,Expected:%d,Got:%d", k, v.Params, v.Result, res)
		}
	}
}

func getTestArr509() []params509 {
	testP1 := params509{
		Params: 2,
		Result: 1,
	}

	testP2 := params509{
		Params: 3,
		Result: 2,
	}

	testP3 := params509{
		Params: 4,
		Result: 3,
	}

	return []params509{testP1, testP2, testP3}
}
