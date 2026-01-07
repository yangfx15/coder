package main

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例 1：
//
// 输入：grid = [
// [1,3,1],
// [1,5,1],
// [4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m][n]
}
