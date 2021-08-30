package main

import "fmt"

func main() {
	r := searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println(r)
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

//二分搜索 >=target的最小值
