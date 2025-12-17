package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var dirs = [][]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}
