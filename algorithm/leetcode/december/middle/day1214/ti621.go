package day1214

/*
本题目巧解法 桶的思想
优先用数量多的去排列
*/

func leastInterval(tasks []byte, n int) int {
	max := 0
	mp := map[byte]int{}
	for _, v := range tasks {
		mp[v]++
	}
	for _, v := range mp {
		if v >= max {
			max = v
		}
	}
	result := (max - 1) * (n + 1)
	for _, v := range mp {
		if v == max {
			result++
		}
	}
	//返回 len(tasks)和 result 的最大值
	if len(tasks) > result {
		return len(tasks)
	}
	return result
}
