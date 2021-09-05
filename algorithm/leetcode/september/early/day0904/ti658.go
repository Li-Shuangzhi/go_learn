package main

import (
	"fmt"
	"sort"
)

func main() {
	res := findClosestElements1([]int{1, 2, 3, 4, 5}, 4, -1)
	fmt.Println(res)
}

func findClosestElements1(arr []int, k int, x int) (res []int) {
	for k > 0 {
		index := sort.SearchInts(arr, x)
		if index == 0 {
			res = append(res, arr[0])
			arr = arr[1:]
		} else if index == len(arr) {
			res = append(res, arr[len(arr)-1])
			arr = arr[:len(arr)-1]
		} else {
			tmp := 0
			if abs(arr[index]-x) >= abs(arr[index-1]-x) {
				tmp = index - 1
			} else {
				tmp = index
			}
			res = append(res, arr[tmp])
			arr = append(arr[:tmp], arr[tmp+1:]...)
		}
		k--
	}
	sort.Ints(res)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	start, end := 0, len(arr)-k
	for start < end {
		mid := start + (end-start)/2
		if x-arr[mid] > arr[mid+k]-x {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return arr[start : start+k]
}
