package day1101

func distributeCandies(candyType []int) int {
	mp := make(map[int]int)
	for _, v := range candyType {
		mp[v]++
	}
	totalType := len(mp)
	oneChildAmount := len(candyType) / 2
	return min(oneChildAmount, totalType)
}

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}
