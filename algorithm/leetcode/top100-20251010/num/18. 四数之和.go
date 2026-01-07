package main

import "sort"

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
//
// 示例 1：
//
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 示例 2：
//
// 输入：nums = [2,2,2,2,2], target = 8
// 输出：[[2,2,2,2]]
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		curNum := target - nums[i]
		curList := threeSum(nums[i+1:], curNum)
		for _, list := range curList {
			res = append(res, append(list, nums[i]))
		}
	}
	return res
}

func threeSum(nums []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		curNum := target - nums[i]
		for l, r := i+1, len(nums)-1; l < r; {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if nums[l]+nums[r] == curNum {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			}
			if nums[l]+nums[r] > curNum {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
