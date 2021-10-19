package hot100

import (
	"algo/algo/hot100"
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	type params struct {
		Params1 string
		Params2 int
		Result  string
	}

	testP1 := params{
		Params1: "PAYPALISHIRING",
		Params2: 3,
		Result:  "PAHNAPLSIIGYIR",
	}

	testP2 := params{
		Params1: "PAYPALISHIRING",
		Params2: 4,
		Result:  "PINALSIGYAHRPI",
	}

	testP3 := params{
		Params1: "A",
		Params2: 1,
		Result:  "A",
	}

	testArr := []params{testP1, testP2, testP3}
	for k, v := range testArr {
		res := hot100.Convert(v.Params1, v.Params2)
		if res != v.Result {
			t.Errorf("Case:%d,params1: %s, params2:%d,Expected:%s,Got:%s", k, v.Params1, v.Params2, v.Result, res)
			break
		}
	}
}

func TestSingle(t *testing.T) {
	type params struct {
		Params1 string
		Params2 int
		Result  string
	}

	testP1 := params{
		Params1: "PAYPALISHIRING",
		Params2: 3,
		Result:  "PAHNAPLSIIGYIR",
	}

	res := hot100.Convert(testP1.Params1, testP1.Params2)
	if res != testP1.Result {
		fmt.Print("err")
	}
}
