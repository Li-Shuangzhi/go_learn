package main

import "fmt"

func main() {
	res := search([]int{3, 1}, 1)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left+1)/2
		if nums[mid] == target {
			return mid
		}
		//left 到 late 有序
		if nums[mid] > nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
