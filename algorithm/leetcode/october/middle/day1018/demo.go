package main

import "fmt"

func main() {
	a := [][]int{
		{-2, -8, -9, 8},
		{-10, 0, 6, -8},
		{-10, -6, 6, 9},
	}
	b := [][]int{
		{4, -7, 5, -5, 9},
		{10, -2, -10, 5, 5},
		{-3, -7, -3, 8, -2},
		{-6, 7, 7, 3, -2},
	}
	res := matrixMultiply(a, b)
	fmt.Println(res)
}

func matrixMultiply(a [][]int, b [][]int) [][]int {
	size1 := len(a)
	size2 := len(a[0])
	size3 := len(b)
	size4 := len(b[0])
	if size3 != size2 {
		return nil
	}
	var c [][]int
	c = make([][]int, size1)
	for i := 0; i < len(c); i++ {
		c[i] = make([]int, size4)
	}

	for i := 0; i < size1; i++ {
		for j := 0; j < size4; j++ {
			sum := 0
			for k := 0; k < size2; k++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}
	return c
}
