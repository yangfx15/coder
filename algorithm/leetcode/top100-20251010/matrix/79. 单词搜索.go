package main

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 示例 1：
//
// 输入：board = [
// ['A','B','C','E'],
// ['S','F','C','S'],
// ['A','D','E','E']], word = "ABCCED"
// 输出：true
// 示例 2：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"
// 输出：true
// 示例 3：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"
// 输出：false
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	read := make([][]bool, m)
	for i := 0; i < m; i++ {
		read[i] = make([]bool, n)
	}

	var dfs func(cur string, i, j int, visited [][]bool) bool
	dfs = func(cur string, i, j int, visited [][]bool) bool {
		if cur == "" {
			return true
		}
		visited[i][j] = true
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			if board[x][y] != cur[0] || visited[x][y] {
				continue
			}
			if dfs(cur[1:], x, y, visited) {
				return true
			}
		}
		visited[i][j] = false
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}

			if dfs(word[1:], i, j, read) {
				return true
			}
		}
	}

	return false
}
