package day1115

import (
	"math"
	"sort"
)

var candidateList []int
var result [][]int
var ways []int

func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	ways = make([]int, 0)
	sort.Ints(candidates)
	candidateList = candidates
	backTrace(target, 0, 0)
	return result
}

func backTrace(target int, sum int, startIndex int) {
	if sum > target {
		return
	}
	if target == sum {
		tmp := make([]int, len(ways))
		copy(tmp, ways)
		result = append(result, tmp)
		return
	}
	set := make(map[int]bool)
	for i := startIndex; i < len(candidateList); i++ {
		if set[candidateList[i]] {
			continue
		} else {
			set[candidateList[i]] = true
		}
		ways = append(ways, candidateList[i])
		backTrace(target, sum+candidateList[i], i+1)
		ways = ways[:len(ways)-1]
	}
}
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}
