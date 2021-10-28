package hot100

import (
	"algo/algo/common"
	"algo/algo/hot100"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	testArr := getTestArr101()
	for k, v := range testArr {
		res := hot100.IsSymmetric(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d, params:%+v, Expected:%+v, Got:%+v", k, v.Params, v.Result, res)
		}
	}
}

type params101 struct {
	Params *common.TreeNode
	Result bool
}

func getTestArr101() []params101 {
	n1 := common.NewTreeNode(1)

	n1.SetRChild(common.NewTreeNode(2))
	n1.SetLChild(common.NewTreeNode(2))

	n1.Left.SetRChild(common.NewTreeNode(4))
	n1.Left.SetLChild(common.NewTreeNode(3))

	n1.Right.SetRChild(common.NewTreeNode(3))
	n1.Right.SetLChild(common.NewTreeNode(4))
	testP1 := params101{
		Params: n1,
		Result: true,
	}

	n2 := common.NewTreeNode(1)
	n2.SetLChild(common.NewTreeNode(2))
	n2.SetRChild(common.NewTreeNode(2))
	n2.Left.SetRChild(common.NewTreeNode(3))
	n2.Right.SetRChild(common.NewTreeNode(3))
	testP2 := params101{
		Params: n2,
		Result: false,
	}

	return []params101{testP1, testP2}
}
