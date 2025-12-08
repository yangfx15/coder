package main

import "fmt"

//给你一个字符串 s，找到 s 中最长的 回文 子串。
//
//
//
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"
//
//
//提示：
//
//1 <= s.length <= 1000

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var res string
	var b, e int
	for i := 0; i < len(s); i++ {
		b1, e1 := getPalindrome(s, i, i)
		b2, e2 := getPalindrome(s, i, i+1)
		if e1-b1 > e2-b2 {
			b, e = b1, e1
		} else {
			b, e = b2, e2
		}
		if e-b > len(res) {
			res = s[b:e]
		}
	}
	return res
}

func getPalindrome(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i, j = i-1, j+1
		//fmt.Printf("i, j: %v, %v\n", i, j)
	}
	//fmt.Printf("i, j: %v, %v\n", i, j)
	return i+1, j
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("cdddc"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("bb"))
}
