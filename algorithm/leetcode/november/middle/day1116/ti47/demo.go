package ti47

var result [][]int
var path []int
var gNums []int

func permuteUnique(nums []int) [][]int {
	result = make([][]int, 0)
	path = make([]int, 0)
	used := make([]bool, len(nums))
	gNums = nums
	backTrace(used)
	return result
}

func backTrace(used []bool) {
	if len(path) == len(gNums) {
		tmp := make([]int, len(gNums))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}
	set := make(map[int]bool)
	for i := 0; i < len(gNums); i++ {
		if !used[i] {
			if set[gNums[i]] {
				continue
			}
			set[gNums[i]] = true
			path = append(path, gNums[i])
			used[i] = true
			backTrace(used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}
