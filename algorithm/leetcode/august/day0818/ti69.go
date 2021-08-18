package main

import "fmt"

func main() {
	sqrt := mySqrt(16)
	fmt.Println(sqrt)
}

/*
实现int sqrt(int x)函数。
计算并返回x的平方根，其中x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/
func mySqrt1(x int) int {
	i := 0
	for true {
		if i*i <= x && (i+1)*(i+1) > x {
			return i
		}
		i++
	}
	return 0
}

//二分查找
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left, right := 0, x
	for left <= right {
		mid := (right-left)/2 + left
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return 0
}
