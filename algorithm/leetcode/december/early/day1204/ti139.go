package day1204

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n)
	set := make(map[string]bool)
	for _, v := range wordDict {
		set[v] = true
	}
	for i := 1; i <= n; i++ {
		if set[s[:i]] {
			dp[i-1] = true
		} else {
			for j := 0; j < i; j++ {
				if dp[j] && set[s[j+1:i]] {
					dp[i-1] = true
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}
