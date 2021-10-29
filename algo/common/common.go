package common

// 单向链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

//根据数组返回单向链表
func NewListNode(arr []int) *ListNode {
	var head ListNode
	var pre ListNode
	for _, num := range arr {
		node := ListNode{Val: num, Next: nil}
		if head.Next == nil {
			head.Next = &node
		}
		if pre.Next == nil {
			pre.Next = &node
		} else {
			pre.Next.Next = &node
			pre = *pre.Next
		}
	}

	return head.Next
}

//根据数组返回单向链表
func NewCycleListNode(arr []int, pos int) *ListNode {
	var head ListNode
	var pre ListNode
	arrNode := []*ListNode{}

	for _, num := range arr {
		node := ListNode{Val: num, Next: nil}
		arrNode = append(arrNode, &node)

		if head.Next == nil {
			head.Next = &node
		}

		if pre.Next == nil {
			pre.Next = &node
		} else {
			pre.Next.Next = &node
			pre = *pre.Next
		}
	}

	if pos != -1 {
		arrNode[len(arr)-1].Next = arrNode[pos]
	}

	return head.Next
}

//二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTreeNode(n int) *TreeNode {
	return &TreeNode{Val: n}
}

func (t *TreeNode) SetRChild(n *TreeNode) {
	t.Right = n
}

func (t *TreeNode) SetLChild(n *TreeNode) {
	t.Left = n
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
