package test

import (
	"algo/algo"
	"reflect"
	"testing"
)

type AddTwoNumParams struct {
	Params1 *algo.ListNode
	Params2 *algo.ListNode
	Result  []int
}

func TestAddTwoNum(t *testing.T) {
	//输入：l1 = [2,4,3], l2 = [5,6,4] 输出：[7,0,8]
	l1_1 := &algo.ListNode{
		Val:  3,
		Next: nil,
	}

	l1_2 := &algo.ListNode{
		Val:  4,
		Next: l1_1,
	}

	l1 := &algo.ListNode{
		Val:  2,
		Next: l1_2,
	}

	l2_2 := &algo.ListNode{
		Val:  4,
		Next: nil,
	}

	l2_1 := &algo.ListNode{
		Val:  6,
		Next: l2_2,
	}

	l2 := &algo.ListNode{
		Val:  5,
		Next: l2_1,
	}

	//输入：l1 = [0], l2 = [0] 输出：[0]
	l3 := &algo.ListNode{
		Val: 0,
	}

	l4 := &algo.ListNode{
		Val: 0,
	}

	//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9] 输出：[8,9,9,9,0,0,0,1]
	l5_1 := &algo.ListNode{
		Val:  9,
		Next: nil,
	}

	l5_2 := &algo.ListNode{
		Val:  9,
		Next: l5_1,
	}

	l5_3 := &algo.ListNode{
		Val:  9,
		Next: l5_2,
	}

	l5_4 := &algo.ListNode{
		Val:  9,
		Next: l5_3,
	}

	l5_5 := &algo.ListNode{
		Val:  9,
		Next: l5_4,
	}

	l5_6 := &algo.ListNode{
		Val:  9,
		Next: l5_5,
	}

	l5 := &algo.ListNode{
		Val:  9,
		Next: l5_6,
	}

	l6_1 := &algo.ListNode{
		Val:  9,
		Next: nil,
	}

	l6_2 := &algo.ListNode{
		Val:  9,
		Next: l6_1,
	}

	l6_3 := &algo.ListNode{
		Val:  9,
		Next: l6_2,
	}

	l6 := &algo.ListNode{
		Val:  9,
		Next: l6_3,
	}

	params := []AddTwoNumParams{
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

	for k, v := range params {
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
