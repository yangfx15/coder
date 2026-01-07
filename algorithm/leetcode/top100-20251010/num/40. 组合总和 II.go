package main

import (
	"fmt"
	"sort"
)

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
// 示例 1:
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
// 示例 2:
//
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
func combinationSum2(candidates []int, target int) [][]int {
	// 排序
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] <= candidates[j]
	})

	var findTarget func(idx, tg int, list []int)
	var res [][]int
	findTarget = func(idx, tg int, list []int) {
		if tg == 0 {
			temp := make([]int, len(list))
			copy(temp, list)
			res = append(res, temp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[idx] > tg {
				break
			}
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			list = append(list, candidates[i])
			findTarget(i+1, tg-candidates[i], list)
			list = list[:len(list)-1]
		}
	}

	findTarget(0, target, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	//fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
