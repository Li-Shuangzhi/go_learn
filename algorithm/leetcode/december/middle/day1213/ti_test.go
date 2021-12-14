package day1213

import (
	"fmt"
	"sort"
	"testing"
)

func TestTi(t *testing.T) {
	nums := []int{1, 0, 1}
	moveZeroes(nums)
	t.Log(nums)
}

func TestLg(t *testing.T) {

	data := []int{10, 11, 17}
	i := sort.Search(len(data), func(i int) bool {
		return data[i] >= 23
	})
	fmt.Println("最终的结果为", i)
}
