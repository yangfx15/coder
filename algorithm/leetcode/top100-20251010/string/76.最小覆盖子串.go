package main

import "fmt"

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
// 示例 2：
//
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
// 示例 3:
//
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 提示：
//
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s 和 t 由英文字母组成
//
// 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
func minWindow(s string, t string) string {
	// 边界情况处理
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// 滑动窗口
	l, r := 0, 0

	// 记录t长度，当前最小子串长度，及返回子串值
	size, curSize, curStr := len(t), 0, ""

	// 记录需覆盖子串元素，及当前子串元素
	record, cur := make(map[uint8]int, 0), make(map[uint8]int, 0)
	for i := 0; i < size; i++ {
		record[t[i]]++
	}

	for r < len(s) {
		idx := s[r]
		r++

		// 当元素存在t中，计数+1
		if record[idx] > cur[idx] {
			curSize++
		}
		cur[idx]++

		// 当前子串包含t里所有元素时
		for l <= r && curSize == size {
			if curStr == "" || r-l < len(curStr) {
				curStr = s[l:r]
			}
			reduce := s[l]

			l++
			cur[reduce]--
			if cur[reduce] < record[reduce] {
				curSize--
			}
		}
	}

	return curStr
}

func main() {
	//fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	//fmt.Println(minWindow("a", "a"))
	//fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("ab", "b"))
}
