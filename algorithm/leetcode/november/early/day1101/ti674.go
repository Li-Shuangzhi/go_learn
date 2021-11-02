package day1101

import (
	"fmt"
)

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 0
	for i := 1; i < len(dp); i++ {
		if dp[i] <= dp[i-1] {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + 1
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	fmt.Println(dp)
	return res
}
