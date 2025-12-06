package main

import "fmt"

// 峰值元素是指其值严格大于左右相邻值的元素。
//
//给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
//你可以假设 nums[-1] = nums[n] = -∞ 。
//
//你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
//
//示例 1：
//
//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引 2。
//示例 2：
//
//输入：nums = [1,2,1,3,5,6,4]
//输出：1 或 5
//解释：你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。
//
//提示：
//
//1 <= nums.length <= 1000
//-231 <= nums[i] <= 231 - 1
//对于所有有效的 i 都有 nums[i] != nums[i + 1]

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// 新增边界值，忽略边界处理
	newNums := []int{-1<<31 - 1}
	newNums = append(newNums, nums...)
	newNums = append(newNums, -1<<31-1)

	// 注意l,r取值
	l, r := 1, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if newNums[mid] > newNums[mid-1] && newNums[mid] > newNums[mid+1] {
			return mid - 1
		}
		// 假设峰值在偏左位置
		if newNums[mid-1] > newNums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 理论上不可能出现
	return -1
}

func main() {
	//fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	//fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	//fmt.Println(findPeakElement([]int{1}))
	//fmt.Println(findPeakElement([]int{1, 5}))
	fmt.Println(findPeakElement([]int{-2147483647, -2147483648}))
	fmt.Println(findPeakElement([]int{1, 5, 4}))
}
