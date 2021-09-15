package main

import (
	"fmt"
)

func main() {
	res := checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()")
	fmt.Println(res)
}

func checkValidString1(s string) bool {
	left, star := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left > 0 {
				left--
			} else if star > 0 {
				star--
			} else {
				return false
			}
		} else {
			star++
		}
	}
	fmt.Println(left, star)
	return left <= star
}

func checkValidString(s string) bool {
	var leftStack, starStack []int
	for i, ch := range s {
		if ch == '(' {
			leftStack = append(leftStack, i)
		} else if ch == '*' {
			starStack = append(starStack, i)
		} else {
			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1]
			} else if len(starStack) > 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		}
	}
	i := len(leftStack) - 1
	j := len(starStack) - 1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if starStack[j] < leftStack[i] {
			return false
		}
	}
	return i < 0
}
