package main

import "fmt"

func main() {
	judge := findJudge1(1, nil)
	fmt.Println(judge)
}

func findJudge1(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	mp := make(map[int]int)
	for i := 0; i < len(trust); i++ {
		mp[trust[i][1]]++
	}
	var trustedByOthers []int
	for i, v := range mp {
		if v == n-1 {
			trustedByOthers = append(trustedByOthers, i)
		}
	}
	if len(trustedByOthers) < 1 {
		return -1
	}
	flags := make([]bool, len(trustedByOthers))
	for j, v := range trustedByOthers {
		for i := 0; i < len(trust); i++ {
			if trust[i][0] == v {
				flags[j] = true
			}
		}
	}
	for i := 0; i < len(flags); i++ {
		if !flags[i] {
			return trustedByOthers[i]
		}
	}
	return -1
}

func findJudge(n int, trust [][]int) int {
	inToValues := make([]int, n+1)
	for _, v := range trust {
		inToValues[v[0]]--
		inToValues[v[1]]++
	}
	for i := 1; i <= n; i++ {
		if inToValues[i] == n-1 {
			return i
		}
	}
	return -1
}
