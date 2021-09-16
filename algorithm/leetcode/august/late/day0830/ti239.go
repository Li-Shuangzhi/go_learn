package main

import (
	"fmt"
	"sort"
)

func main() {
	res := maxSlidingWindow([]int{1, -1}, 1)
	fmt.Println(res)
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if n <= k {
		return []int{sum}
	}
	result := make([]int, n-k+1)
	slide := make([]int, k)
	copy(slide, nums[:k])
	sort.Ints(slide)
	result[0] = slide[len(slide)-1]
	for i := 1; i <= n-k; i++ {
		slide = binaryDelete(nums[i-1], slide)
		slide = binaryAdd(nums[i+k-1], slide)
		result[i] = slide[len(slide)-1]
	}
	return result
}

//二分 插入
func binaryAdd(num int, slide []int) []int {
	if len(slide)-1 < 0 || num >= slide[len(slide)-1] {
		res := make([]int, len(slide))
		copy(res, slide)
		res = append(res, num)
		return res
	}
	left, right := 0, len(slide)-1
	for left < right {
		mid := left + (right-left)/2
		if slide[mid] < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	tmp := slide[left:]
	res := make([]int, len(slide))
	copy(res, slide)
	res = append(res[:left], num)
	res = append(res, tmp...)
	return res
}

//二分 删除
func binaryDelete(num int, slide []int) []int {
	left, right := 0, len(slide)-1
	for left <= right {
		mid := left + (right-left)/2
		if slide[mid] > num {
			right = mid - 1
		} else if slide[mid] < num {
			left = mid + 1
		} else {
			res := make([]int, len(slide))
			copy(res, slide)
			res = append(res[:mid], res[mid+1:]...)
			return res
		}
	}
	return nil
}
