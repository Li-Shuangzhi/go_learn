package day1101

import "fmt"

func longestPalindromeSubseq(s string) int {
	size := len(s)
	//init array dp
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}
	for j := 0; j < size; j++ {
		for i := size - 1; i >= 0; i-- {
			if i <= j {
				if i == j {
					dp[i][j] = 1
				} else if j-i == 1 {
					if s[i] == s[j] {
						dp[i][j] = 2
					} else {
						dp[i][j] = 1
					}
				} else {
					t := 0
					for k := i; k < j; k++ {
						if s[k] == s[j] {
							if k+1 == j {
								t = 2
							} else {
								t = 2 + dp[k+1][j-1]
							}
							break
						}
					}
					dp[i][j] = max(t, dp[i][j-1])
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[0][size-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
