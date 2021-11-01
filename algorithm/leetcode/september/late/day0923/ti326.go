package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree1(81))
}

func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}
	return n%3 == 0 && isPowerOfThree1(n/3)
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}
	return n%3 == 0 && isPowerOfThree(n/3)
}
