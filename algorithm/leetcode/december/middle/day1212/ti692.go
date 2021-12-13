package day1212

import "sort"

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int)
	for _, value := range words {
		mp[value]++
	}
	uniqueWords := make([]string, 0, len(mp))
	for w := range mp {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return mp[s] > mp[t] || mp[s] == mp[t] && s < t
	})
	return uniqueWords[:k]
}
