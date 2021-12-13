package main

import "fmt"

var m, n int
var horseX, horseY int

func main() {
	_, err := fmt.Scanf("%d %d %d %d", &m, &n, &horseX, &horseY)
	if err != nil {
		return
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		if check(0, i) {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i <= m; i++ {
		if check(i, 0) {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if check(i, j) {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp[m][n])
}

func check(x, y int) bool {
	if (abs(x-horseX) == 1 && abs(y-horseY) == 2) || (abs(x-horseX) == 2 && abs(y-horseY) == 1) {
		return true
	}
	if x == horseX && y == horseY {
		return true
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
