package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	myAtoi("")
}

func myAtoi(s string) int {
	//字符串->字符数组 便于操作
	trimed := strings.TrimSpace(s)
	res := 0
	end := 0
	for i := len(trimed) - 1; i >= 0; i-- {
		v := trimed[i]
		if !(v >= '0' && v <= '9') {
			end = i
		}
	}
	if end != 0 {
		trimed = trimed[:end]
	}
	fmt.Println(trimed)
	bytes := []byte(trimed)
	sign := true
	for _, v := range bytes {
		res *= 10
		if v >= '0' && v <= '9' {
			n := v - '0'
			res += int(n)
		} else if v == '-' {
			sign = false
		} else if v == '+' {
			continue
		} else {
			return 0
		}
	}
	if !sign {
		res = -res
	}
	if res < math.MinInt32 {
		return math.MinInt32
	} else if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
