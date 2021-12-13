package day1109

import "fmt"

/*
	907. 子数组的最小值之和
	给定一个整数数组 arr，找到 min(b)的总和，其中 b 的范围为 arr 的每个（连续）子数组。
	由于答案可能很大，因此 返回答案模 10^9 + 7 。
*/

func sumSubarrayMins(arr []int) (ans int) {
	stack := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
		ans += stack[0]
		fmt.Println(stack, ans)
	}
	return ans
}

//连续
