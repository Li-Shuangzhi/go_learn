package main

import (
	"fmt"
	"sort"
)

func main() {

	res := rangeBitwiseAnd(5, 7)
	fmt.Println(res)
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[n][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func arrangeCoins1(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}

func arrangeCoins(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left+1)/2
		if mid*(mid+1) <= 2*n {
			left = mid
		} else {
			right = mid - 1
		}
	}
	fmt.Println(left, right)
	return left
}

func rangeBitwiseAnd(left int, right int) int {
	result := left
	for i := left + 1; i <= right; i++ {
		result = result & i
	}
	return result
}
