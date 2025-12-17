package main

import "fmt"

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [
// ["1","0","1","0","0"],
// ["1","0","1","1","1"],
// ["1","1","1","1","1"],
// ["1","0","0","1","0"]]
// 输出：4
// 示例 2：
//
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1
// 示例 3：
//
// 输入：matrix = [["0"]]
// 输出：0
func maximalSquare(matrix [][]byte) int {
	// 动态规划
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '1' {
				continue
			}
			dp[i][j] = 1
			if i != 0 && j != 0 {
				continue
			}
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + 1
			}
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}

	var maxArea int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '1' {
				continue
			}
			dp[i][j] = 1
			if i > 0 && j > 0 && dp[i-1][j] >= 1 && dp[i][j-1] >= 1 && dp[i-1][j-1] >= 1 {
				minV := min(dp[i-1][j], dp[i][j-1])
				minV = min(dp[i-1][j-1], minV)
				dp[i][j] = minV + 1
			}
			maxArea = max(maxArea, dp[i][j]*dp[i][j])
		}
	}
	//fmt.Println(dp)
	return maxArea
}

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		{'0', '1'},
		{'1', '0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '1', '1', '1'},
	}))
}
