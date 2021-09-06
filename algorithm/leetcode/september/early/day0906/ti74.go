package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 4, 7, 9, 23, 29, 33}
	x := 34
	i := sort.Search(len(data), func(i int) bool { return data[i] > x })
	fmt.Println(i)
	//if i < len(data) && data[i] == x {
	//	// x is present at data[i]
	//} else {
	//	// x is not present in data,
	//	// but i is the index where it would be inserted.
	//}

	res := searchMatrix([][]int{{1}}, 1)
	fmt.Println(res)
}

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	row := sort.Search(n, func(i int) bool {
		return matrix[i][len(matrix[i])-1] >= target
	})
	if row >= n {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
