package hot100

import (
	"algo/algo/common"
)

/*
94. 二叉树的中序遍历:https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

示例 4：
输入：root = [1,2]
输出：[2,1]

示例 5：
输入：root = [1,null,2]
输出：[1,2]


提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶: 递归算法很简单，你可以通过迭代算法完成吗？

中序遍历：首先遍历左子树，然后访问根结点，最后遍历右子树

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 注意：此处[]int 不好做append，直接使用匿名函数做全局的[]int
func InorderTraversal(root *common.TreeNode) (res []int) {
	var travel func(node *common.TreeNode)

	travel = func(node *common.TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)
		res = append(res, node.Val)

		travel(node.Right)
	}

	travel(root)

	return
}
