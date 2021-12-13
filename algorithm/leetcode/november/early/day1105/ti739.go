package day1105

type Pair struct {
	index       int
	temperature int
}

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	ans := make([]int, len(temperatures))
	stack = append(stack, 0)
	for i := 1; i < len(temperatures); i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
