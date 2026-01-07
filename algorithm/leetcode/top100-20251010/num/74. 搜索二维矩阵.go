package main

// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：matrix = [
// [1,3,5,7],
// [10,11,16,20],
// [23,30,34,60]], target = 3
// 输出：true
// 示例 2：
//
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出：false
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	up, down := 0, len(matrix)-1
	for up <= down {
		m1 := (up + down) / 2
		nums := matrix[m1]
		if nums[0] <= target && target <= nums[len(nums)-1] {
			return searchSlice(nums, target)
		}
		if target < nums[0] {
			down = m1 - 1
		}
		if target > nums[len(nums)-1] {
			up = m1 + 1
		}
	}
	return false
}

func searchSlice(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
