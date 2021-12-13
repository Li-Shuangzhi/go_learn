package day1115

import (
	"sort"
)

func longestWord(words []string) string {
	set := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		set[words[i]] = true
	}
	var list []string
	for i := 0; i < len(words); i++ {
		list = append(list, words[i])
		for j := 1; j < len(words[i]); j++ {
			if !set[words[i][:j]] {
				list = list[:len(list)-1]
				break
			}
		}
	}
	sort.Strings(list)
	maxLen := 0
	for i := 0; i < len(list); i++ {
		if len(list[i]) > maxLen {
			maxLen = len(list[i])
		}
	}
	for i := 0; i < len(list); i++ {
		if len(list[i]) == maxLen {
			return list[i]
		}
	}
	return ""
}
