package day1213

func moveZeroes(nums []int) {
	n := len(nums)
	i, j := 0, 0
	for j < n && i < n {
		for i < n && nums[i] != 0 {
			i++
		}
		j = i
		for j < n && nums[j] == 0 {
			j++
		}
		if j >= n || i >= n || i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}
