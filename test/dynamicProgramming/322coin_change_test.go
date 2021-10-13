package dynamicProgramming

import (
	"algo/algo/dynamicProgramming"
	"testing"
)

/*func TestCoinChange(t *testing.T) {
	testArr := getTestArr322()
	for k, v := range testArr {
		res := dynamicProgramming.CoinChange(v.Params1, v.Params2)
		if res != v.Result {
			t.Errorf("Case:%d,params1: %+v, params2: %d,Expected:%d,Got:%d", k, v.Params1, v.Params2, v.Result, res)
		}
	}
}*/

func TestCoinChangeDp(t *testing.T) {
	testArr := getTestArr322()
	for k, v := range testArr {
		res := dynamicProgramming.CoinChangeDp(v.Params1, v.Params2)
		if res != v.Result {
			t.Errorf("Case:%d,params1: %+v, params2: %d,Expected:%d,Got:%d", k, v.Params1, v.Params2, v.Result, res)
		}
	}
}

type params322 struct {
	Params1 []int
	Params2 int
	Result  int
}

func getTestArr322() []params322 {
	testP1 := params322{
		Params1: []int{1, 2, 5},
		Params2: 11,
		Result:  3,
	}

	testP2 := params322{
		Params1: []int{2},
		Params2: 3,
		Result:  -1,
	}

	testP3 := params322{
		Params1: []int{1},
		Params2: 0,
		Result:  0,
	}

	testP4 := params322{
		Params1: []int{1},
		Params2: 1,
		Result:  1,
	}

	testP5 := params322{
		Params1: []int{1},
		Params2: 2,
		Result:  2,
	}

	return []params322{testP1, testP2, testP3, testP4, testP5}
}
