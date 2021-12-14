package day1213

func findAnagrams(s string, p string) []int {
	n := len(s)
	m := len(p)
	if m > n {
		return nil
	}
	pmap := make(map[byte]int)
	for i := 0; i < m; i++ {
		pmap[p[i]]++
	}
	tmap := make(map[byte]int)
	for i := 0; i < m; i++ {
		tmap[s[i]]++
	}
	var res []int
	for i := 0; i <= n-m; i++ {
		if check(pmap, tmap) {
			res = append(res, i)
		}
		if i != n-m {
			tmap[s[i]]--
			tmap[s[i+m]]++
		}
	}
	return res
}

func check(pmap, tmap map[byte]int) bool {
	for k, v := range pmap {
		if tmap[k] != v {
			return false
		}
	}
	return true
}
