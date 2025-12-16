package main

import "fmt"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
func generateParenthesis(n int) []string {
	var res []string
	var generate func(cur string, l, r int)
	generate = func(cur string, l, r int) {
		if r == 0 && l == 0 {
			res = append(res, cur)
			return
		}
		if l > r {
			return
		}
		if l > 0 {
			generate(cur+"(", l-1, r)
		}
		if r > 0 {
			generate(cur+")", l, r-1)
		}
	}
	generate("", n, n)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
}
