package day1211

import (
	"math"
	"sort"
)

var tmp []int
var Nums []int
var mp map[int][]int
var maxKey int

func maxSubsequence1(nums []int, k int) []int {
	maxKey = math.MaxInt64
	mp = make(map[int][]int)
	Nums = nums
	backTrace(0, k)
	max := math.MinInt32
	for key, _ := range mp {
		if key > max {
			max = key
		}
	}
	return mp[max]
}

func backTrace(start int, k int) {
	if len(tmp) == k {
		sum := 0
		for i := 0; i < k; i++ {
			sum += tmp[i]
		}
		if sum < maxKey {
			return
		}
		t := make([]int, k)
		copy(t, tmp)
		mp[sum] = t
		if sum > maxKey {
			maxKey = sum
		}
		return
	}
	for i := start; i < len(Nums); i++ {
		tmp = append(tmp, Nums[i])
		backTrace(i+1, k)
		tmp = tmp[:len(tmp)-1]
	}
}

type Pair struct {
	index int
	value int
}

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	pairs := make([]*Pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = &Pair{
			index: i,
			value: nums[i],
		}
	}
	sort.Slice(pairs, func(a, b int) bool {
		if pairs[a].value > pairs[b].value {
			return true
		}
		return false
	})
	tmp := make([]*Pair, k)
	for i := 0; i < k; i++ {
		tmp[i] = pairs[i]
	}
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].index < tmp[j].index {
			return true
		}
		return false
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = tmp[i].value
	}
	return res
}
