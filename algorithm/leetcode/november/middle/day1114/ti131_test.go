package day1114

import (
	"fmt"
	"testing"
)

type Student1 struct {
}

func TestTi(t *testing.T) {
	r := restoreIpAddresses("1111")
	fmt.Println(len(r))
	fmt.Println(r)
}

func TestArray(t *testing.T) {
	var nums [10]*Student1
	fmt.Println(nums)
	fmt.Println(nums[0] == nil)
}
