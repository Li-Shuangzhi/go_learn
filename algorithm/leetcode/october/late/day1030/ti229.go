package main

import "fmt"

func main() {
	res := majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2})
	fmt.Println(res)
}

func majorityElement(nums []int) (res []int) {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	for k, v := range mp {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return
}
