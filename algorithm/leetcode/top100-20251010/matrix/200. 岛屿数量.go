package main

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
//
//示例 1：
//
//输入：grid = [
//  ['1','1','1','1','0'],
//  ['1','1','0','1','0'],
//  ['1','1','0','0','0'],
//  ['0','0','0','0','0']
//]
//输出：1
//示例 2：
//
//输入：grid = [
//  ['1','1','0','0','0'],
//  ['1','1','0','0','0'],
//  ['0','0','1','0','0'],
//  ['0','0','0','1','1']
//]
//输出：3

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '0'
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= len(grid) || y >= len(grid[0]) || x<0 || y<0 {
				continue
			}
			if grid[x][y] == '0' {
				continue
			}
			dfs(x, y)
		}
	}

	var num int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			num++
			dfs(i, j)
		}
	}

	return num
}
