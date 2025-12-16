package main

import "fmt"

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
// 示例 2：
//
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
// 示例 3：
//
// 输入：nums = [1,0,1,2]
// 输出：3
//
// 提示：
//
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
func longestConsecutive(nums []int) int {
	rd := make(map[int]bool, 0)
	for _, num := range nums {
		rd[num] = true
	}
	var maxV int
	for _, num := range nums {
		if rd[num-1] {
			continue
		}
		curNum := num + 1
		for rd[curNum] {
			curNum ++
		}
		if curNum-num > maxV {
			maxV = curNum - num
		}
	}
	return maxV
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
}
