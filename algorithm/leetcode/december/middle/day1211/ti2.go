package day1211

// 前缀数组
// decr[i]表示i天为连续下降的第几天
// incr[i]表示i天为连续上升的第几天
// 满足decr[i]>=time and incr[i]>=time的加入结果集

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	res := make([]int, 0)
	decr := make([]int, n)
	incr := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			decr[i] = decr[i-1] + 1
		} else {
			decr[i] = 0
		}
	}
	for i := n - 2; i > 0; i-- {
		if security[i] <= security[i+1] {
			incr[i] = incr[i+1] + 1
		} else {
			incr[i] = 0
		}
	}
	for i := time; i < n-time; i++ {
		if incr[i] >= time && decr[i] >= time {
			res = append(res, i)
		}
	}
	return res
}
