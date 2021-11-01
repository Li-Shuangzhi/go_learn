package main

import "fmt"

func main() {
	res := countMaxOrSubsets([]int{2, 2, 2})
	fmt.Println(res)
}

var max = 0
var amount = 0

func countMaxOrSubsets(nums []int) int {
	max, amount = 0, 0
	vector := make([]int, 0)
	dfs(nums, 0, vector)
	return amount
}

func dfs(nums []int, start int, vector []int) {
	i := 0
	for i = start; i < len(nums); i++ {
		vector = append(vector, nums[i])
		dfs(nums, i+1, vector)
		vector = vector[:len(vector)-1]
	}
	if i == len(nums) {
		res := 0
		for j := 0; j < len(vector); j++ {
			res |= vector[j]
		}
		if res > max {
			max = res
			amount = 1
		} else if res == max {
			amount++
		}
		return
	}
}
