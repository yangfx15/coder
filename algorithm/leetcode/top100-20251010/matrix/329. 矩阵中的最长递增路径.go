package main

//给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
//
//示例 1：
//
//输入：matrix = [
//[9,9,4],
//[6,6,8],
//[2,1,1]]
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

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	var getLongest func(i, j int) int
	getLongest = func(i, j int) int {
		maxV := 1
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if matrix[i][j] < matrix[x][y] {
				cur := dp[x][y]
				if cur == 0 {
					cur = getLongest(x, y)
				}

				if maxV < cur+1 {
					maxV = cur + 1
				}
			}
		}
		dp[i][j] = maxV
		return maxV
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, getLongest(i, j))
		}
	}
	return res
}
