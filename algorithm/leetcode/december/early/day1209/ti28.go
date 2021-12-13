package day1209

func strStr1(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if 0 == n {
		return 0
	}
	if m < n {
		return -1
	}
	for i := range haystack {
		k := i
		j := 0
		flag := true
		for j = range needle {
			if k >= m || haystack[k] != needle[j] {
				flag = false
				break
			}
			k++
		}
		if flag {
			return i
		}
	}
	return -1
}

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	for i := 0; i <= n-m; i++ {
		flag := true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
