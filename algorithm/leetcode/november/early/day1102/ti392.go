package day1102

//isSubsequence1 double pointer
func isSubsequence1(s string, t string) bool {
	pointer1, pointer2 := 0, 0
	for pointer1 < len(s) && pointer2 < len(t) {
		if s[pointer1] == t[pointer2] {
			pointer1++
			pointer2++
		} else {
			pointer2++
		}
	}
	return pointer1 == len(s)
}

//isSubsequence dynamic program
func isSubsequence(s, t string) bool {
	if len(s) < 1 {
		return true
	}
	sLength, tLength := len(s), len(t)
	dp := make([][]int, sLength+1)
	for i := 0; i <= sLength; i++ {
		dp[i] = make([]int, tLength+1)
	}
	for i := 1; i <= sLength; i++ {
		for j := 1; j <= tLength; j++ {
			if s[i-1] == t[j-1] {
				for k := 0; k < j; k++ {
					dp[i][j] = max(dp[i][j], dp[i-1][k]+1)
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i <= tLength; i++ {
		if dp[len(s)][i] == len(s) {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
