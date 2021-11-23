package hot100

import (
	"algo/algo/common"
)

/*
543. 二叉树的直径:https://leetcode-cn.com/problems/diameter-of-binary-tree/
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树
          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
*/

var m int

func DiameterOfBinaryTree(root *common.TreeNode) int {
	return depth(root)
}

func depth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	L := depth(root.Left)
	R := depth(root.Right)
	m = max(L+R+1, m)

	return max(L, R) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
