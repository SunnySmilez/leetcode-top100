package hot100

import "algo/algo/common"

/*
234. 回文链表:https://leetcode-cn.com/problems/palindrome-linked-list/
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
示例 1：
输入：head = [1,2,2,1]
输出：true

示例 2：
输入：head = [1,2]
输出：false

思路1：将链表值存入数组中，使用双指针遍历
思路2：使用快慢指针找到链表中间点，翻转后半部分链表，再比对链表
思路3：进行递归，使用递归第一个返回是最后一个值的特性进行比较
*/

// 思路1：将链表值存入数组中，使用双指针遍历
func IsPalindrome234(head *common.ListNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i := 0; i < len(nums); i++ {
		l := i
		r := len(nums) - 1 - i

		if nums[l] != nums[r] {
			return false
		}
	}

	return true
}

// 思路2：将链表后半部分进行翻转，使用快慢指针遍历
func IsPalindrome2342(head *common.ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	halfNode := endOfHalf(head)
	reverseNode := reverse1(halfNode.Next)
	for reverseNode != nil {
		if head.Val != reverseNode.Val {
			return false
		}

		head = head.Next
		reverseNode = reverseNode.Next
	}

	return true
}

func reverse1(head *common.ListNode) *common.ListNode {
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

func endOfHalf(head *common.ListNode) *common.ListNode {
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/*
func IsPalindrome2341(head *common.ListNode) bool {


	return
}

func dg(node *common.ListNode) int {
	for node.Next == nil {
		return node.Val
	}

	return dg(node)
}*/
