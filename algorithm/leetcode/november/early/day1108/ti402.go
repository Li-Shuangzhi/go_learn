package day1108

import "strings"

func removeKdigits(num string, k int) string {
	//维护一个单调栈(递增)
	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if "" == ans {
		ans = "0"
	}
	return ans
}
