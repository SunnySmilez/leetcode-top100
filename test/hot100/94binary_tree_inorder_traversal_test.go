package hot100

import (
	"algo/algo/common"
	"algo/algo/hot100"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testArr := getTestArr94()
	for k, v := range testArr {
		res := hot100.InorderTraversal(v.Params)
		if reflect.DeepEqual(res, v.Result) == false {
			t.Errorf("Case:%d, params:%+v, Expected:%+v, Got:%+v", k, v.Params, v.Result, res)
		}
	}
}

type params94 struct {
	Params *common.TreeNode
	Result []int
}

func getTestArr94() []params94 {
	n1 := common.NewTreeNode(1)
	n1.SetRChild(common.NewTreeNode(2))
	n1.Right.SetLChild(common.NewTreeNode(3))
	testP1 := params94{
		Params: n1,
		Result: []int{1, 3, 2},
	}

	n2 := common.NewTreeNode(1)
	n2.SetLChild(common.NewTreeNode(2))
	testP2 := params94{
		Params: n2,
		Result: []int{2, 1},
	}

	n3 := common.NewTreeNode(1)
	n3.SetRChild(common.NewTreeNode(2))
	testP3 := params94{
		Params: n3,
		Result: []int{1, 2},
	}

	return []params94{testP1, testP2, testP3}
}
