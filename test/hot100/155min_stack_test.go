package hot100

import (
	"algo/algo/hot100"
	"testing"
)

func Test155MinStack(t *testing.T) {
	a := hot100.Constructor()
	a.Push(-2)
	a.Push(0)
	a.Push(-3)
	num1 := a.GetMin()
	if num1 != -3 {
		t.Errorf("num1 expect -3 got %d", num1)
	}

	a.Pop()
	num2 := a.Top()
	if num2 != 0 {
		t.Errorf("num2 expect 0 got %d", num2)
	}

	num3 := a.GetMin()
	if num3 != -2 {
		t.Errorf("num3 expect -2 got %d", num3)
	}
}

func Test155MinStack1(t *testing.T) {
	a := hot100.Constructor()
	a.Push(2)
	a.Push(0)
	a.Push(3)
	a.Push(0)
	num1 := a.GetMin()
	if num1 != 0 {
		t.Errorf("num1 expect 0 got %d", num1)
	}

	a.Pop()
	num2 := a.GetMin()
	if num2 != 0 {
		t.Errorf("num2 expect 0 got %d", num1)
	}

	a.Pop()
	num3 := a.GetMin()
	if num3 != 0 {
		t.Errorf("num3 expect 0 got %d", num3)
	}

	a.Pop()
	num4 := a.GetMin()
	if num4 != 2 {
		t.Errorf("num4 expect 2 got %d", num1)
	}
}
