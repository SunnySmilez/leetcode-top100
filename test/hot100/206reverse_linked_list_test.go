package hot100

import (
	"algo/algo/common"
	"algo/algo/hot100"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	testArr := getTestArr206()
	for k, v := range testArr {
		res := hot100.HasCycle1(v.Params)
		if reflect.DeepEqual(res, v.Result) {
			t.Errorf("Case:%d, params:%+v, Expected:%+v, Got:%+v", k, v.Params, v.Result, res)
		}
	}
}

type params206 struct {
	Params *common.ListNode
	Result *common.ListNode
}

func getTestArr206() []params206 {
	testP1 := params206{
		Params: common.NewListNode([]int{1, 2, 3, 4, 5}),
		Result: common.NewListNode([]int{5, 4, 3, 2, 1}),
	}

	testP2 := params206{
		Params: common.NewListNode([]int{1, 2}),
		Result: common.NewListNode([]int{2, 1}),
	}

	testP3 := params206{
		Params: common.NewListNode([]int{}),
		Result: common.NewListNode([]int{}),
	}

	return []params206{testP1, testP2, testP3}
}
