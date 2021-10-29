package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testArr := getTestArr136()
	for k, v := range testArr {
		res := hot100.SingleNumber(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d, params:%+v, Expected:%d, Got:%d", k, v.Params, v.Result, res)
		}
	}
}

type params136 struct {
	Params []int
	Result int
}

func getTestArr136() []params136 {
	testP1 := params136{
		Params: []int{2, 2, 1},
		Result: 1,
	}

	testP2 := params136{
		Params: []int{4, 1, 2, 1, 2},
		Result: 4,
	}

	return []params136{testP1, testP2}
}
