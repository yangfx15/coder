package main

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
	if len(s) <= 1 {
		return 0
	}

	var stack []int
	var res, prev int
	for i := 0; i < len(s); i++ {
		if s[i] == ')' && len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur := i-pop+1+prev
			res = max(res, cur)
			prev = cur
		} else {
			prev = 0
		}
		if s[i] == '(' {
			stack = append(stack, i)
		}
	}

	return res
}
