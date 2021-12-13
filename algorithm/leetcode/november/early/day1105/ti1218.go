package day1105

import "log"

func longestSubsequence1(arr []int, difference int) (ans int) {
	size := len(arr)
	mp := make(map[int]int)
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 1
	}
	for i, v := range arr {
		t := v - difference
		if v1, ok := mp[t]; ok {
			dp[i] = dp[v1] + 1
		}
		mp[v] = i
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	log.Println(dp)
	return ans
}
