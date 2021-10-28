package hot100

import (
	"algo/algo/common"
	"algo/algo/hot100"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testArr := getTestArr104()
	for k, v := range testArr {
		res := hot100.MaxDepth1(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d, params:%+v, Expected:%+v, Got:%+v", k, v.Params, v.Result, res)
		}
	}
}

type params104 struct {
	Params *common.TreeNode
	Result int
}

func getTestArr104() []params104 {
	n1 := common.NewTreeNode(3)

	n1.SetLChild(common.NewTreeNode(9))
	n1.SetRChild(common.NewTreeNode(20))

	n1.Right.SetRChild(common.NewTreeNode(7))
	n1.Right.SetLChild(common.NewTreeNode(15))
	testP1 := params104{
		Params: n1,
		Result: 3,
	}

	return []params104{testP1}
}
