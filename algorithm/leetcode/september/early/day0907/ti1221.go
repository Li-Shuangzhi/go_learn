package main

import "fmt"

func main() {
	res := balancedStringSplit("R")
	fmt.Println(res)
}

func balancedStringSplit(s string) (ans int) {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
		if mp['R'] == mp['L'] {
			mp['R'] = 0
			mp['L'] = 0
			ans++
		}
	}
	return
}
