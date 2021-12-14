package day1213

import (
	"math"
	"strings"
)

/*
函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。

*/
func myAtoi(s string) int {
	flag := true
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	i := 0
	if s[i] == '-' {
		flag = false
		s = s[1:]
	} else if s[i] == '+' {
		flag = true
		s = s[1:]
	} else if !(s[i] >= '0' && s[i] <= '9') {
		return 0
	}
	s = strings.TrimLeft(s, "0")
	//去除后置非法字符
	s = strings.TrimRightFunc(s, func(r rune) bool {
		if !(r >= '0' && r <= '9') {
			return true
		}
		return false
	})
	var res int64
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res *= 10
			res += int64(c - '0')
			var t int64
			if !flag {
				t = -res
			} else {
				t = res
			}
			if t > math.MaxInt32 {
				return math.MaxInt32
			}
			if t < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	if !flag {
		res *= -1
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	return int(res)
}
