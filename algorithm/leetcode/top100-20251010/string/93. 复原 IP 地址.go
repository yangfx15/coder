package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
// 示例 1：
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：
//
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
func restoreIpAddresses(s string) []string {
	var res []string
	var generate func(i int, ip []string)

	generate = func(idx int, ip []string) {
		if idx == len(s) && len(ip) == 4 {
			res = append(res, strings.Join(ip, "."))
			return
		}

		for i := 1; i <= 3; i++ {
			if idx+i > len(s) || (s[idx] == '0' && i > 1) {
				break
			}
			cur := s[idx : idx+i]
			num, _ := strconv.Atoi(cur)
			if num < 0 || num > 255 {
				break
			}
			ip = append(ip, cur)
			generate(idx+i, ip)
			ip = ip[:len(ip)-1]
		}
	}
	generate(0, []string{})
	return res
}

func main()  {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}