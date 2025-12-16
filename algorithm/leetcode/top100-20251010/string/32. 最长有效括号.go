package main

import "fmt"

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串 的长度。
//
//左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 "(()())"。
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0

func longestValidParentheses(s string) int {
	stack := []int{-1}
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
				continue
			}
			pop := stack[len(stack)-1]
			res = max(res, i-pop)
		} else {
			stack = append(stack, i)
		}
	}
	return res
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("(()())"))
	fmt.Println(longestValidParentheses("(()(()))"))
	fmt.Println(longestValidParentheses(""))
}
