package day1116

import "sort"

var result [][]int
var path []int

func subsetsWithDup(nums []int) [][]int {
	result = make([][]int, 0)
	path = make([]int, 0)
	sort.Ints(nums)
	backTrace(0, nums, false)
	return result
}

func backTrace(index int, nums []int, used bool) {
	if index == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}
	if index > 0 && !used && nums[index] == nums[index-1] {
		backTrace(index+1, nums, false)
	} else {
		path = append(path, nums[index])
		backTrace(index+1, nums, true)
		path = path[:len(path)-1]
		backTrace(index+1, nums, false)
	}
}
