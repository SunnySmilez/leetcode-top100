package hot100

import (
	"algo/algo/hot100"
	"testing"
)

/*
示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false

示例 5：
输入：s = "{[]}"
输出：true
*/
func TestIsValid(t *testing.T) {
	type params struct {
		Params string
		Result bool
	}

	testP1 := params{
		Params: "()",
		Result: true,
	}

	testP2 := params{
		Params: "()[]{}",
		Result: true,
	}

	testP3 := params{
		Params: "(]",
		Result: false,
	}

	testP4 := params{
		Params: "([)]",
		Result: false,
	}

	testP5 := params{
		Params: "{[]}",
		Result: true,
	}

	testP6 := params{
		Params: "(([]){})",
		Result: true,
	}

	testP7 := params{
		Params: "(){}}{",
		Result: false,
	}

	testArr := []params{testP1, testP2, testP3, testP4, testP5, testP6, testP7}
	for k, v := range testArr {
		res := hot100.IsValid(v.Params)
		if res != v.Result {
			t.Errorf("Case:%d, params: %s, Expected:%t, Got:%t", k, v.Params, v.Result, res)
		}
	}
}
