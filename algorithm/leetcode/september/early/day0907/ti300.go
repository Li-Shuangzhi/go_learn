package main

import (
	"fmt"
)

func main() {
	res := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	ans := 0
	dp[0] = 1
	for i := 0; i < n; i++ {
		tmp := nums[i]
		for j := 0; j < i; j++ {
			if nums[j] < tmp {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	fmt.Println(dp)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
