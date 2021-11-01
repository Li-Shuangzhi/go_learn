package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTwoEvents([][]int{
		{1, 3, 2},
		{4, 5, 2},
		{2, 4, 3},
	}))
}

//maxTwoEvents
func maxTwoEvents(events [][]int) (res int) {
	sort.Slice(events, func(x, y int) bool {
		return events[x][1] < events[y][1]
	})
	fmt.Println(events)

	//f[i] 前i项中的最大价值
	f := make([]int, len(events))
	f[0] = events[0][2]
	for i := 1; i < len(events); i++ {
		f[i] = max(f[i-1], events[i][2])
	}
	fmt.Println(f)
	for _, v := range events {
		//在终止时间list里找比v[0]小的最大值
		left, right := 0, len(events)-1
		for left < right {
			mid := left + (right-left)/2 + 1
			if events[mid][1] < v[0] {
				left = mid
			} else {
				right = mid - 1
			}
		}
		tmp := v[2]
		if v[0] > events[left][1] {
			tmp += f[left]
		}
		res = max(res, tmp)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
