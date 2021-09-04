package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 5, 1, 9, 0, 10, 6, 13, 23, 16}
	sort.Ints(nums)
	res := binarySearch(nums, 23)
	fmt.Println(res)
}

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
