package main

import "fmt"

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
//
// 示例 1：
//
// 输入：num1 = "11", num2 = "123"
// 输出："134"
// 示例 2：
//
// 输入：num1 = "456", num2 = "77"
// 输出："533"
// 示例 3：
//
// 输入：num1 = "0", num2 = "0"
// 输出："0"
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	var add int
	var res string
	for ; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		cur := add
		if i >= 0 {
			cur += int(num1[i] - '0')
		}
		if j >= 0 {
			cur += int(num2[j] - '0')
		}
		res = fmt.Sprintf("%d%s", cur%10, res)
		add = cur/10
	}
	return res
}

func main()  {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
}