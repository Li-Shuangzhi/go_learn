package main

func main() {

}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = s[i] == s[j] && (i-1 < j+1 || dp[i-1][j+1])
			if dp[i][j] && i-j+1 > len(ans) {
				ans = s[j : i+1]
			}
		}
	}
	return ans
}
