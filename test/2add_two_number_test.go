package test

import (
	"algo/algo"
	"algo/algo/common"
	"reflect"
	"testing"
)

func TestAddTwoNum(t *testing.T) {
	type params struct {
		Params1 *common.ListNode
		Params2 *common.ListNode
		Result  []int
	}

	//输入：l1 = [2,4,3], l2 = [5,6,4] 输出：[7,0,8]
	l1 := common.NewListNode([]int{2, 4, 3})
	l2 := common.NewListNode([]int{5, 6, 4})

	//输入：l1 = [0], l2 = [0] 输出：[0]
	l3 := common.NewListNode([]int{0})
	l4 := common.NewListNode([]int{0})

	//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9] 输出：[8,9,9,9,0,0,0,1]
	l5 := common.NewListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l6 := common.NewListNode([]int{9, 9, 9, 9})

	testArr := []params{
		{ //输入：l1 = [2,4,3], l2 = [5,6,4] 输出：[7,0,8]
			Params1: l1,
			Params2: l2,
			Result:  []int{7, 0, 8},
		},
		{ //输入：l1 = [0], l2 = [0] 输出：[0]
			Params1: l3,
			Params2: l4,
			Result:  []int{0},
		},
		{ //输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9] 输出：[8,9,9,9,0,0,0,1]
			Params1: l5,
			Params2: l6,
			Result:  []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for k, v := range testArr {
		res := algo.AddTwoNumbers(v.Params1, v.Params2)
		resSlice := make([]int, len(v.Result), cap(v.Result))
		i := 0
		for res != nil {
			resSlice[i] = res.Val
			res = res.Next
			i++
		}

		if reflect.DeepEqual(v.Result, resSlice) == false {
			t.Error("Case:", k, "Error, Expected:", v.Result, ",Got:", resSlice)
		}
	}
}
