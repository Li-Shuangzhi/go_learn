package main

import "fmt"

func main() {
	res := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	fmt.Println(res)
}

//寻找旋转排序数组中的最小值
//[3,4,5,1,2]
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
