package hot100

import (
	"algo/algo/common"
)

/*
21. 合并两个有序链表：https://leetcode-cn.com/problems/merge-two-sorted-lists
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

func MergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	head := &common.ListNode{}
	node := head // 这个写法很好，不需要再下面再去给哑结点赋值
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}

	return head.Next
}
