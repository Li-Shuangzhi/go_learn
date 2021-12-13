package day1105

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//单调栈
	mp := make(map[int]int)
	stack := make([]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		for len(stack) > 0 && num > nums2[stack[len(stack)-1]] {
			mp[nums2[stack[len(stack)-1]]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	//fmt.Println(mp)
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if v, ok := mp[nums1[i]]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}
