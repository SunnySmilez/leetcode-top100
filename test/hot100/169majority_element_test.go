package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testArr := getTestArr169()
	for k, v := range testArr {
		res := hot100.MajorityElement1(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d, params:%+v, Expected:%d, Got:%d", k, v.Params, v.Result, res)
		}
	}
}

type params169 struct {
	Params []int
	Result int
}

func getTestArr169() []params169 {
	testP1 := params169{
		Params: []int{3, 2, 3},
		Result: 3,
	}

	testP2 := params169{
		Params: []int{2, 2, 1, 1, 1, 2, 2},
		Result: 2,
	}

	return []params169{testP1, testP2}
}
