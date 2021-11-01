package main

import (
	"fmt"
	"strings"
)

func main() {
	res := lengthOfLastWord1("   fly me   to   the moon  ")
	r := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(res, r)
}

//"   fly me   to   the moon  "
func lengthOfLastWord1(s string) int {
	size := len(s)
	left, right := -1, -1
	for i := size - 1; i >= 0; i-- {
		if !(s[i] == ' ') && right == -1 {
			right = i
		}
		if right != -1 && (i == 0 || s[i] == ' ') {
			if i == 0 && s[0] != ' ' {
				left = -1
				break
			}
			left = i
			break
		}
	}
	fmt.Println(right, left)
	return right - left
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	splits := strings.Split(s, " ")
	fmt.Println(splits)
	size := len(splits)
	return len(splits[size-1])
}

func lengthOfLastWord2(s string) int {
	s = strings.Trim(s, " ")
	return len(s[strings.LastIndex(s, " ")+1:])
}
