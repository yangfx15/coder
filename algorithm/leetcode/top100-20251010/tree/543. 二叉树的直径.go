package main

// 给你一棵二叉树的根节点，返回该树的 直径 。
//
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
//
// 两节点之间路径的 长度 由它们之间边数表示。
//
// 示例 1：
//
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
// 示例 2：
//
// 输入：root = [1,2]
// 输出：1
func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(cur *TreeNode, lv int) int
	var res int
	dfs = func(cur *TreeNode, lv int) int {
		if cur == nil {
			return 0
		}
		l := dfs(cur.Left, lv+1)
		r := dfs(cur.Right, lv+1)

		res = max(res, l+r)
		return max(r, l) + 1
	}
	dfs(root, 0)

	return res
}
