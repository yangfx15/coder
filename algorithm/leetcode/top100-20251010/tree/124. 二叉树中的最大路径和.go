package main

import (
	"fmt"
	"math"
)

// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。
// 该路径 至少包含一个 节点，且不一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
// 示例 1：
//
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
// 示例 2：
//
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var dfs func(cur *TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}
		if cur.Left == nil && cur.Right == nil {
			maxSum = max(maxSum, cur.Val)
			return cur.Val
		}

		curMax := cur.Val
		l, r := dfs(cur.Left), dfs(cur.Right)
		if l > 0 {
			curMax += l
		}
		if r > 0 {
			curMax += r
		}
		maxSum = max(maxSum, curMax)

		return max(cur.Val, cur.Val+max(l, r))
	}

	maxSum = max(maxSum, dfs(root))
	return maxSum
}

func main() {
	fmt.Println(maxPathSum(&TreeNode{Val: 2, Right: &TreeNode{Val: -1}}))
	fmt.Println(maxPathSum(&TreeNode{Val: -2, Right: &TreeNode{Val: -1}}))
}
