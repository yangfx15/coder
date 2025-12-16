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
	record, courTotal, total := make(map[int]int, 0), make(map[int]int, 0), make(map[int]struct{}, 0)
	for _, cour := range prerequisites {
		c1, c2 := cour[0], cour[1]
		courTotal[c2] = c1
		total[c1], total[c2] = struct{}{}, struct{}{}
		record[c1]++
	}

	var cours []int
	for k := range total {
		// 记录不需要先导课程的
		if record[k] == 0 {
			cours = append(cours, k)
		}
	}

	fmt.Println(record)
	fmt.Println(courTotal)

	learnNum := 0
	for len(cours) > 0 {
		size := len(cours)
		learnNum += size
		fmt.Println(cours)

		for i := 0; i < size; i++ {
			cour := cours[0]
			cours = cours[1:]
			tail := courTotal[cour]
			record[tail]--

			if record[tail] == 0 {
				cours = append(cours, tail)
			}
		}
	}

	return learnNum == len(total)
}

func main() {
	//fmt.Println(canFinish(2, [][]int{{1, 0}}))
	//fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}
