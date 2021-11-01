package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := areNumbersAscending("9 hi 9 integer mom")
	fmt.Println(res)
}

func areNumbersAscending(s string) bool {
	splits := strings.Split(s, " ")
	var t int64 = -1
	for _, v := range splits {
		if v[0] >= '0' && v[0] <= '9' {
			//是数字
			parseInt, _ := strconv.ParseInt(v, 10, 64)
			if t >= parseInt {
				return false
			}
			t = parseInt
		}
	}
	return true
}
