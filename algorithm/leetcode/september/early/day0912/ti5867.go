package main

import "fmt"

func main() {
	prefix := reversePrefix("abcd", 'z')
	fmt.Println(prefix)
}

func reversePrefix(word string, ch byte) string {
	w := []byte(word)
	for i := 0; i < len(w); i++ {
		if word[i] == ch {
			//找到了
			for j := 0; j <= i/2; j++ {
				w[j], w[i-j] = w[i-j], w[j]
			}
			break
		}
	}
	return string(w)
}
