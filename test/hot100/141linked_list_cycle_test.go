package hot100

import (
	"algo/algo/common"
	"algo/algo/hot100"
	"testing"
)

func TestHasCycle(t *testing.T) {
	testArr := getTestArr141()
	for k, v := range testArr {
		res := hot100.HasCycle1(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d, params:%+v, Expected:%t, Got:%t", k, v.Params, v.Result, res)
		}
	}
}

type params141 struct {
	Params *common.ListNode
	Result bool
}

func getTestArr141() []params141 {
	l1 := common.NewCycleListNode([]int{3, 2, 0, -4}, 1)
	testP1 := params141{
		Params: l1,
		Result: true,
	}

	testP2 := params141{
		Params: common.NewCycleListNode([]int{1, 2}, 0),
		Result: true,
	}

	testP3 := params141{
		Params: common.NewCycleListNode([]int{1}, -1),
		Result: false,
	}

	return []params141{testP1, testP2, testP3}
}
