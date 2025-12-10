package main

import "fmt"

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//示例 1：
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//示例 2：
//
//输入：height = [4,2,0,3,2,5]
//输出：9

func trap(height []int) int {
	// 1.定义一个先进后出的栈，用来记录水位的左边界和最低点
	var stack []int

	// 2.定义返回参数
	var area int
	for i := 0; i < len(height); i++ {
		// 3.当前高度高于栈内最后一个数时，代表可以存储雨水，开始循环处理
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			remain := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 4.若当前栈内没有元素，表示缺少左边界，根据木桶效应，这时没法存储雨水
			if len(stack) == 0 {
				continue
			}
			pre := stack[len(stack)-1]
			area += (min(height[i], height[pre]) - height[remain]) * (i - pre - 1)

			// 调试信息
			//fmt.Printf("area:%v, i:%v, pre:%v, min:%v, remain:%v, size:%v \n", area, i, pre, min(height[i], height[pre]), height[remain], i-pre-1)
		}

		stack = append(stack, i)
	}

	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
