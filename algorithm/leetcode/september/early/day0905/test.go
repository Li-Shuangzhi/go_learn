package main

import "fmt"

func main() {
	fmt.Println(myPow(34.00515, -3))
}
func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1.0 / x
	}
	t := myPow(x, n/2)
	if n%2 != 0 {
		return t * t * myPow(x, n%2)
	} else {
		return t * t
	}
}

func fib(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	c := 0
	for n > 1 {
		c = a + b
		a, b = b, c
		n--
	}
	return c % (1e9 + 7)
}

func findMiddleIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	tsum := 0
	for i := 0; i < len(nums); i++ {
		if tsum == sum-tsum-nums[i] {
			return i
		} else {
			tsum += nums[i]
		}
	}
	return -1
}
