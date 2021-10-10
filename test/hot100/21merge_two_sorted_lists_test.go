package hot100

import (
	"algo/algo/common"
	"algo/algo/hot100"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type Params struct {
		Params1 *common.ListNode
		Params2 *common.ListNode
		Result  *common.ListNode
	}

	//输入：l1 = [1,2,4], l2 = [1,3,4] 输出：[1,1,2,3,4,4]
	t1params1 := common.NewListNode([]int{1, 2, 4})
	t1params2 := common.NewListNode([]int{1, 3, 4})
	t1result1 := common.NewListNode([]int{1, 1, 2, 3, 4, 4})

	testP1 := Params{
		Params1: t1params1,
		Params2: t1params2,
		Result:  t1result1,
	}

	//输入：l1 = [], l2 = [] 输出：[]
	t2params1 := common.NewListNode([]int{})
	t2params2 := common.NewListNode([]int{})
	t2result1 := common.NewListNode([]int{})
	testP2 := Params{
		Params1: t2params1,
		Params2: t2params2,
		Result:  t2result1,
	}

	//输入：l1 = [], l2 = [0] 输出：[0]
	t3params1 := common.NewListNode([]int{})
	t3params2 := common.NewListNode([]int{0})
	t3result1 := common.NewListNode([]int{0})
	testP3 := Params{
		Params1: t3params1,
		Params2: t3params2,
		Result:  t3result1,
	}

	testArr := []Params{testP1, testP2, testP3}

	for k, v := range testArr {
		res := hot100.MergeTwoLists(v.Params1, v.Params2)
		if reflect.DeepEqual(res, v.Result) == false {
			t.Errorf("Case:%d, params1: %+v, params2: %+v, Expected:%+v, Got:%+v", k, v.Params1, v.Params2, v.Result, res)
		}
	}
}
