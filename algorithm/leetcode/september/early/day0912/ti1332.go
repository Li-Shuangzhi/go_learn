package main

func main() {

}

func removePalindromeSub(s string) int {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return 2
		}
	}
	return 1
}
