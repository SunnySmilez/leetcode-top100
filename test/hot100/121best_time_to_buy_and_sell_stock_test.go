package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testArr := getTestArr121()
	for k, v := range testArr {
		res := hot100.MaxProfitDp(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d, Expected:%d,Got:%d", k, v.Params, v.Result, res)
		}
	}
}

type params121 struct {
	Params []int
	Result int
}

func getTestArr121() []params121 {
	testP1 := params121{
		Params: []int{7, 1, 5, 3, 6, 4},
		Result: 5,
	}

	testP2 := params121{
		Params: []int{7, 6, 4, 3, 1},
		Result: 0,
	}

	return []params121{testP1, testP2}
}
