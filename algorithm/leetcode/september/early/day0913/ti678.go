package main

import (
	"fmt"
	"strings"
)

func main() {
	res := checkValidString("(*))")
	fmt.Println(res)
}

func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}
	s1 := strings.ReplaceAll(s, "*", ")")
	s2 := strings.ReplaceAll(s, "*", "")
	s3 := strings.ReplaceAll(s, "*", "(")
	return judgeValid(s1) || judgeValid(s2) || judgeValid(s3)
}

func judgeValid(s string) bool {
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1] == s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
