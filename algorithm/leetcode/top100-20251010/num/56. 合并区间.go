package main

import (
	"fmt"
	"sort"
)

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2：
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//示例 3：
//
//输入：intervals = [[4,7],[1,4]]
//输出：[[1,7]]
//解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 先按第一个元素排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i<len(intervals);i ++ {
		val := intervals[i]

		// 没有突破前一个元素的边界，就忽略
		if val[1] <= res[len(res)-1][1] {
			continue
		}

		// 如果有重叠元素，就合并
		if val[0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = val[1]
			continue
		}

		res = append(res, val)
	}
	return res
}

func main()  {
	fmt.Println(merge([][]int{
		{1,3},{2,6},{8,10},{15,18},
	}))

	fmt.Println(merge([][]int{
		{1,4}, {4,5},
	}))

	fmt.Println(merge([][]int{
		{4,7}, {1,4},
	}))
}