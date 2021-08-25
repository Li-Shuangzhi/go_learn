package main

func main() {

}

func longestCommonPrefix(strs []string) (result string) {
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		flag := false
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				flag = true
				break
			}
			if strs[j][i] != c {
				flag = true
				break
			}
		}
		if flag {
			break
		} else {
			result += string(c)
		}
	}
	return
}
