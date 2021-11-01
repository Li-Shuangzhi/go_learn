package main

import (
	"fmt"
	"strings"
)

func main() {
	res := convert1("A", 1)
	fmt.Println(res)
}

func convert1(s string, numRows int) (res string) {
	if numRows == 1 {
		return s
	}
	width := getWidth(s, numRows)
	bytes := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		bytes[i] = make([]byte, width)
	}
	for index := 0; index < len(s); index++ {
		x, y := getPosition(index+1, numRows)
		fmt.Println(x, y)
		bytes[x][y] = s[index]
	}

	for i := 0; i < len(bytes); i++ {
		for j := 0; j < len(bytes[i]); j++ {
			if bytes[i][j] != 0 {
				res += string(bytes[i][j])
			}
		}
	}
	return
}

func getWidth(s string, numRows int) int {
	size := len(s)
	t1 := size / (numRows*2 - 2)
	t2 := size % (numRows*2 - 2)
	k := (numRows - 1) * t1
	if t2 != 0 {
		if t2 <= numRows {
			k += 1
		} else {
			k += 1
			k += t2 - numRows
		}
	}
	return k
}

func getPosition(index int, numRow int) (int, int) {
	t1 := index / (numRow*2 - 2)
	t2 := index % (numRow*2 - 2)
	k := (numRow - 1) * t1
	if t2 != 0 {
		if t2 <= numRow {
			k += 1
		} else {
			k += 1
			k += t2 - numRow
		}
	}
	x := -1
	if t2 == 0 {
		x = 1
	} else if t2 <= numRow && t2 >= 1 {
		x = t2 - 1
	} else {
		x = numRow - 1
		x -= t2 - numRow
	}
	return x, k - 1
}

func convert(s string, numRows int) (res string) {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	n := 2*numRows - 2
	for i, ch := range s {
		x := i % n
		rows[min(x, n-x)] += string(ch)
	}
	return strings.Join(rows, "")
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
