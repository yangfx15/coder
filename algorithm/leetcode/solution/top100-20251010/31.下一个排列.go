package main

import "fmt"

// 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须 原地 修改，只允许使用额外常数空间。
//
//示例 1：
//* 输入：nums = [1,2,3]
//* 输出：[1,3,2]
//
//示例 2：
//* 输入：nums = [3,2,1]
//* 输出：[1,2,3]
//
//示例 3：
//* 输入：nums = [1,1,5]
//* 输出：[1,5,1]
//
//示例 4：
//* 输入：nums = [1]
//* 输出：[1]

// 1234 -> 1243 -> 1324 -> 1342 -> 1423 ...
// 4321 -> 1234

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	// 首先，从个位数开始找到一个比高位数大的数
	firstSmaller := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			firstSmaller = i - 1
			break
		}
	}

	// 再从个位数开始找，找到第一个比当前数大的数，并交换两个数的值
	// 如：1243 -> 1342
	if firstSmaller >= 0 {
		for i := len(nums) - 1; i > firstSmaller; i-- {
			if nums[i] > nums[firstSmaller] {
				nums[i], nums[firstSmaller] = nums[firstSmaller], nums[i]
				break
			}
		}
	}

	// 从当前位置的下一个位置，往后交换位置
	// 如：1342 -> 1324
	for l, r := firstSmaller+1, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	fmt.Println(nums)
}

func main() {
	nextPermutation([]int{1, 2, 3, 4})
	nextPermutation([]int{1, 2, 4, 3})
	nextPermutation([]int{1, 3, 2, 4})
	nextPermutation([]int{4, 3, 2, 1})
	nextPermutation([]int{1, 3, 2})
}
