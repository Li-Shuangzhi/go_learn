package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		set := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if set[s[j]] {
				res = max(res, len(set))
				break
			} else {
				set[s[j]] = true
				res = max(res, len(set))
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
