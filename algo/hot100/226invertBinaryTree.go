package hot100

import "algo/algo/common"

/*
226. 翻转二叉树:https://leetcode-cn.com/problems/invert-binary-tree/
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1


*/
func InvertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return root
	}

	InvertTree(root.Left)
	InvertTree(root.Right)

	r := root.Right
	root.Right = root.Left
	root.Left = r

	return root
}
