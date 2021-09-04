package main

import (
	"fmt"
)

func main() {
	res := findPeakElement([]int{-2147483648, -2147483647})
	fmt.Println(res)
}

func findPeakElement(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		midValue := nums[mid]
		if mid == 0 {
			if midValue > nums[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		} else if mid == n-1 {
			if midValue < nums[mid-1] {
				right = mid - 1
			} else {
				return mid
			}
		} else if midValue > nums[mid-1] && midValue > nums[mid+1] {
			return mid
		} else if midValue < nums[mid-1] && midValue < nums[mid+1] {
			right = mid - 1
		} else if midValue < nums[mid+1] && midValue > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
