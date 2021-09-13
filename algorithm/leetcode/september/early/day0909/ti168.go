package main

import (
	"fmt"
)

func main() {
	res := convertToTitle(28)
	fmt.Println(res)
}

func convertToTitle(columnNumber int) string {
	var ans []byte
	for columnNumber > 0 {
		a0 := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(a0-1))
		columnNumber = (columnNumber - a0) / 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}

func titleToNumber(columnTitle string) (res int) {
	t := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		n := int(columnTitle[i] - 'A' + 1)
		res += t * n
		t *= 26
	}
	return res
}
