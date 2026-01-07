package main

import "fmt"

//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
//在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//
//例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//示例 1：
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//示例 2：
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录课程的先导课程数量 & 总课程数
	preCourse, postCourse := make(map[int]int, 0), make(map[int][]int, 0)
	for _, course := range prerequisites {
		c1, c2 := course[0], course[1]

		// 计算c1的先导课程
		preCourse[c1]++

		// 学完c2学c1
		postCourse[c2] = append(postCourse[c2], c1)
	}

	// 记录无需先导的课程
	couldLearnCour, finished := []int{}, 0
	for i := 0; i < numCourses; i++ {
		if preCourse[i] == 0 {
			couldLearnCour = append(couldLearnCour, i)
			finished++
		}
	}

	// 开始学习
	for len(couldLearnCour) > 0 {
		cur := couldLearnCour[0]
		couldLearnCour = couldLearnCour[1:]

		// 学完当前，就更新后续课程的先导数
		for _, c := range postCourse[cur] {
			preCourse[c]--
			if preCourse[c] <= 0 {
				couldLearnCour = append(couldLearnCour, c)
				finished++
			}
		}
	}

	return finished == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}
