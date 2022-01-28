package hot100

import "algo/algo/common"

/*

19. 删除链表的倒数第 N 个结点:https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]
*/
//1. 遍历获取长度，然后将目标节点指向下下一个节点
//2. 双指针方，两指针间隔n个节点，快指针达到末尾的时候，慢指针即为目标节点
func RemoveNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	return nil
}
