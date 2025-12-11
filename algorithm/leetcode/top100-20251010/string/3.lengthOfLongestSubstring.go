package main

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
	size := len(s)

	// 处理边界Case
	if size == 0 || size == 1 {
		return size
	}

	uniqueKey := make(map[uint8]int, 0)
	longest := 1
	left, right := 0, 0
	// 滑动窗口
	for right = 0; right < size; right++ {
		curKey := s[right]
		uniqueKey[curKey]++

		for uniqueKey[curKey] >= 2 {
			delKey := s[left]
			left++
			uniqueKey[delKey]--
		}
		longest = max(longest, right-left+1)
	}
	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
