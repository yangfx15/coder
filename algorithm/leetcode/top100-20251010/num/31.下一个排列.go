package main

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

// 1.从后往前，找到第一个比尾数小的数，交换两个数
// 2.从当前数往后，找到第一个比尾数大的数，交换所有数
// 1243 -> 1342 -> 1324
// 1342 -> 1432 -> 1423
func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}
	firstSmaller := 0
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			firstSmaller = i
			break
		}
	}
	//fmt.Println("firstSmaller:", firstSmaller)

	for i := n - 1; i > firstSmaller; i-- {
		if nums[i] > nums[firstSmaller] {
			nums[i], nums[firstSmaller] = nums[firstSmaller], nums[i]
			firstSmaller = firstSmaller + 1
			break
		}
	}
	//fmt.Println("firstSmaller:", firstSmaller)

	// 交换
	for i, j := firstSmaller, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	//fmt.Println(nums)
}

func main() {
	nextPermutation([]int{1, 2, 3, 4})
	nextPermutation([]int{1, 2, 4, 3})
	nextPermutation([]int{1, 3, 2, 4})
	nextPermutation([]int{4, 3, 2, 1})
	nextPermutation([]int{1, 3, 2})
	nextPermutation([]int{2, 3, 1})
}
