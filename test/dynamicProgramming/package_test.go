package dynamicProgramming

import (
	"algo/algo/dynamicProgramming"
	"testing"
)

func TestPackage(t *testing.T) {
	testArr := getPackageArr()
	for k, v := range testArr {
		res := dynamicProgramming.Package(v.N, v.W, v.Wt, v.Val)
		if res != v.Result {
			t.Errorf("Case:%d,n:%d,v:%d,wt:%+v,val:%+v,Expected:%d,Got:%d", k, v.N, v.W, v.Wt, v.Val, v.Result, res)
		}
	}
}

type paramsPackage struct {
	N      int
	W      int
	Wt     []int
	Val    []int
	Result int
}

func getPackageArr() []paramsPackage {
	testP1 := paramsPackage{
		N:      3,
		W:      4,
		Wt:     []int{2, 1, 3},
		Val:    []int{4, 2, 3},
		Result: 6,
	}

	return []paramsPackage{testP1}
}
