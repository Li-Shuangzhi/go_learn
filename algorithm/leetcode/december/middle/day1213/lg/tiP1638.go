package main

import "fmt"

func main() {
	m, n := 0, 0
	fmt.Scan(&m, &n)
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&nums[i])
	}
	//以下是算法核心

}
