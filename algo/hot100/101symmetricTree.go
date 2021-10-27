package hot100

import (
	"algo/algo/common"
)

/*
101. 对称二叉树:https://leetcode-cn.com/problems/symmetric-tree/
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/
// 互为镜像可以理解成是两个根节点具有相同的值；同时左子树和右子树互为镜像，也就是左子树的左节点等于右子树的右节点，右子树的左节点等于左子树的右节点，且根节点要相等
func IsSymmetric(root *common.TreeNode) bool {
	return check(root, root)
}

func check(p, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
