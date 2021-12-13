package day1108

import "fmt"

func getHint(secret string, guess string) string {
	size := len(secret)
	var x, y int
	mp1 := make(map[uint8]int)
	mp2 := make(map[uint8]int)
	for i := 0; i < size; i++ {
		if secret[i] == guess[i] {
			x++
		} else {
			mp1[secret[i]]++
			mp2[guess[i]]++
		}
	}
	for k, v := range mp1 {
		if value, has := mp2[k]; has {
			y += min(v, value)
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
