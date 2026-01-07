package main

import "fmt"

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
//
// 提示：
//
// 1 <= k <= nums.length <= 105
// -104 <= nums[i] <= 104
func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	cur := nums[l]
	i, j := l, r
	for i < j {
		for ; nums[i] < cur; i++ {
		}
		for ; nums[j] > cur; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSort(nums, l, j, k)
	}
	return quickSort(nums, j+1, r, k)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	//fmt.Println(findKthLargest([]int{3, 5, 4, 2, 6, 1}, 2))
	//fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
