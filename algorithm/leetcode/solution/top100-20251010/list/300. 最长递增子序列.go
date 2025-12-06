package main

import "fmt"

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
//
//
//示例 1：
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//示例 2：
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//示例 3：
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur > res[len(res)-1] {
			res = append(res, cur)
			continue
		}

		if cur < res[len(res)-1] {
			l, r := 0, len(res)-1
			for l <= r {
				mid := (r + l) / 2
				if cur == res[mid] {
					l = mid
					break
				}
				if cur < res[mid] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			res[l] = cur
		}
		//for j := 0; j < len(res); j++ {
		//	if cur == res[j] {
		//		break
		//	}
		//	if cur < res[j] {
		//		//fmt.Printf("replace:%v->%v, res:%v\n", res[j], cur, res)
		//		res[j] = cur
		//		break
		//	}
		//}
	}
	return len(res)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{1, 2, -10, -8, -7}))
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}
