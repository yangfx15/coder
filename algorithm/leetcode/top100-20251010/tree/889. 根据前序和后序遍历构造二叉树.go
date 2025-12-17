package main

// 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，
// postorder 是同一棵树的后序遍历，重构并返回二叉树。
//
// 如果存在多个答案，您可以返回其中 任何 一个。
//
// 示例 1：
//
// 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
// 输出：[1,2,3,4,5,6,7]
// 示例 2:
//
// 输入: preorder = [1], postorder = [1]
// 输出: [1]
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	root := &TreeNode{Val: preorder[0]}
	idx := 0
	for ; idx < len(postorder); idx++ {
		// 寻找下一个根节点
		if postorder[idx] == preorder[1] {
			break
		}
	}

	// preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
	root.Left = constructFromPrePost(preorder[1:idx+2], postorder[:idx+1])

	// preorder = [1,2], postorder = [2,1]
	root.Right = constructFromPrePost(preorder[idx+2:], postorder[idx+1:len(postorder)-1])

	return root
}
