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
*/
func ReverseList(head *common.ListNode) *common.ListNode {
	var prev *common.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
