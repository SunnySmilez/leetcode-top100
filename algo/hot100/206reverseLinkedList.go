package hot100

import "algo/algo/common"

/*
206. 反转链表:https://leetcode-cn.com/problems/reverse-linked-list/
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
讲解视频：https://leetcode-cn.com/problems/reverse-linked-list/solution/shi-pin-jiang-jie-die-dai-he-di-gui-hen-hswxy/
注意：当遍历到最后一个节点的是，需要做的操作：
	1. 将下一个节点的下一个节点的指针指向当前节点（建立新链表的指针）：node.next.next=node
	2. 将当前节点的下一个节点指向nil(断开就链表的指针)：node.next=nil
*/
func ReverseList(head *common.ListNode) *common.ListNode {
	return reverse(head)
}

func reverse(node *common.ListNode) (newHead *common.ListNode) {
	if node == nil || node.Next == nil {
		return node
	}

	newHead = reverse(node.Next)

	node.Next.Next = node
	node.Next = nil

	return
}

func ReverseList1(head *common.ListNode) *common.ListNode {
	prev := &common.ListNode{}
	curr := head

	for curr.Next == nil {
		next := curr.Next
		curr.Next = prev

		prev = curr.Next
		curr = next
	}

	return curr
}
