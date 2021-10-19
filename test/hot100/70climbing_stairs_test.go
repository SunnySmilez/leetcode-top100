package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testArr := getTestArr70()
	for k, v := range testArr {
		res := hot100.ClimbStairs(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d,params: %d, Expected:%d,Got:%d", k, v.Params, v.Result, res)
		}
	}
}

type params70 struct {
	Params int
	Result int
}

func getTestArr70() []params70 {
	/*testP1 := params70{
		Params:2,
		Result:  2,
	}*/

	testP2 := params70{
		Params: 3,
		Result: 3,
	}

	//return []params70{testP1, testP2}
	return []params70{testP2}
}
