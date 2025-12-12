package main

import "fmt"

//给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
//
//示例 1：
//
//输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
//输出：4
//解释：最长递增路径为 [1, 2, 6, 9]。
//示例 2：
//
//输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
//输出：4
//解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
//示例 3：
//
//输入：matrix = [[1]]
//输出：1

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	dp := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		dp = append(dp, make([]int, len(matrix[i])))
	}

	var maxSize int
	var dfs func(mr [][]int, i, j int) int
	dfs = func(mr [][]int, i, j int) int {
		if dp[i][j] > 0 {
			return dp[i][j]
		}
		dp[i][j] = 1

		// 上下左右遍历
		for _, dir := range dirs {
			a, b := i+dir[0], j+dir[1]
			if a < 0 || b < 0 || a >= len(matrix) || b >= len(matrix[0]) {
				//fmt.Printf("a:%v, b:%v\n", a, b)
				continue
			}

			if mr[a][b] > mr[i][j] {
				dp[i][j] = max(dp[i][j], dfs(mr, a, b)+1)
			}
		}

		return dp[i][j]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxSize = max(maxSize, dfs(matrix, i, j))
		}
	}
	fmt.Println(dp)

	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// [9,9,4],[6,6,8],[2,1,1]
	// [3,4,5],[3,2,6],[2,2,1]
	//fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	//fmt.Println(longestIncreasingPath([][]int{{1, 2}}))
}
