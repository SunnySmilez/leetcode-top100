package dynamicProgramming

import (
	"algo/algo/dynamicProgramming"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	testArr := getTestArr300()
	for k, v := range testArr {
		res := dynamicProgramming.LengthOfLIS(v.Params1)
		if res != v.Result {
			t.Errorf("Case:%d,params1: %+v,Expected:%d,Got:%d", k, v.Params1, v.Result, res)
		}
	}
}

type params300 struct {
	Params1 []int
	Result  int
}

func getTestArr300() []params300 {
	testP1 := params300{
		Params1: []int{10, 9, 2, 5, 3, 7, 101, 18},
		Result:  4,
	}

	testP2 := params300{
		Params1: []int{0, 1, 0, 3, 2, 3},
		Result:  4,
	}

	testP3 := params300{
		Params1: []int{7, 7, 7, 7, 7, 7, 7},
		Result:  1,
	}

	return []params300{testP1, testP2, testP3}
}
