package test

import (
	"algo/algo"
	"testing"
)

/*
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
*/
func TestAddTwoNum1(t *testing.T) {
	node := &algo.ListNode{
		Val:  3,
		Next: nil,
	}

	node2 := &algo.ListNode{
		Val:  4,
		Next: node,
	}

	l1 := &algo.ListNode{
		Val:  2,
		Next: node2,
	}

	n3 := &algo.ListNode{
		Val:  4,
		Next: nil,
	}

	n2 := &algo.ListNode{
		Val:  6,
		Next: n3,
	}

	l2 := &algo.ListNode{
		Val:  5,
		Next: n2,
	}

	res := algo.AddTwoNumbers(l1, l2)
	for res != nil {
		print(res.Val)
		res = res.Next
	}
}

/*
输入：l1 = [0], l2 = [0]
输出：[0]
*/
func TestAddTwoNum2(t *testing.T) {
	l1 := &algo.ListNode{
		Val: 0,
	}

	l2 := &algo.ListNode{
		Val: 0,
	}

	res := algo.AddTwoNumbers(l1, l2)
	for res != nil {
		print(res.Val)
		res = res.Next
	}
}

/*
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/
func TestAddTwoNum3(t *testing.T) {
	node1 := algo.ListNode{
		Val:  9,
		Next: nil,
	}

	node2 := algo.ListNode{
		Val:  9,
		Next: &node1,
	}

	node3 := algo.ListNode{
		Val:  9,
		Next: &node2,
	}

	node4 := algo.ListNode{
		Val:  9,
		Next: &node3,
	}

	node5 := algo.ListNode{
		Val:  9,
		Next: &node4,
	}

	node6 := algo.ListNode{
		Val:  9,
		Next: &node5,
	}

	l1 := algo.ListNode{
		Val:  9,
		Next: &node6,
	}

	n1 := algo.ListNode{
		Val:  9,
		Next: nil,
	}

	n2 := algo.ListNode{
		Val:  9,
		Next: &n1,
	}

	n3 := algo.ListNode{
		Val:  9,
		Next: &n2,
	}

	l2 := algo.ListNode{
		Val:  9,
		Next: &n3,
	}

	res := algo.AddTwoNumbers(&l1, &l2)
	for res != nil {
		print(res.Val)
		res = res.Next
	}
}
