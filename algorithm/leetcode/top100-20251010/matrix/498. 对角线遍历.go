package main

import "fmt"

//给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
//示例 1：
//输入：mat = [
//[1,2,3],
//[4,5,6],
//[7,8,9]]
//输出：[1,2,4,7,5,3,6,8,9]

//示例 2：
//输入：mat = [
//[1,2],
//[3,4]]
//输出：[1,2,3,4]

// 对角线原则：m+n
// 上升：x--, y++ => 起点：i<m(i,0), i>=m(m-1,i-m+1) => min(i,m-1),max(0,i-m+1)
// 下降：y--, x++ => 起点：i<n(0,i), i>=n(i-n+1,n-1) => max(0,i-n+1),min(i,n-1)
// 先上升，后下降
func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}

	m, n := len(mat), len(mat[0])
	res := make([]int, m*n)
	var idx int
	for i := 0; i < (m + n - 1); i++ {
		if i%2 == 0 {
			x, y := min(i, m-1), max(0, i-m+1)
			for x >= 0 && y <= n-1 {
				res[idx] = mat[x][y]
				x--
				y++
				idx++
			}
		} else {
			x, y := max(0, i-n+1), min(i, n-1)
			for y >= 0 && x <= m-1 {
				res[idx] = mat[x][y]
				x++
				y--
				idx++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
