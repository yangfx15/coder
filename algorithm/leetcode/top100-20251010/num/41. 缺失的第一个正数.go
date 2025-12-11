package main

import "fmt"

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//示例 1：
//
//输入：nums = [1,2,0]
//输出：3
//解释：范围 [1,2] 中的数字都在数组中。
//示例 2：
//
//输入：nums = [3,4,-1,1]
//输出：2
//解释：1 在数组中，但 2 没有。
//示例 3：
//
//输入：nums = [7,8,9,11,12]
//输出：1
//解释：最小的正数 1 没有出现。
//
//
//提示：
//
//1 <= nums.length <= 105
//-231 <= nums[i] <= 231 - 1

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 先把所有值置为正数，如果当前数不属于有效索引，就置为数组一个不存在的数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 有效值：正整数1~n，所以当某个位置有值时，就修改该值为0
	for i := 0; i < n; i++ {
		cur := abs(nums[i])
		if cur >= 1 && cur <= n {
			// 把这个数置为负数
			nums[cur-1] = -abs(nums[cur-1])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{1}))
}
