package day1204

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i, j := 0, 1
	for j < n {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}

func removeElement(nums []int, val int) int {
	n := len(nums)
	i, j := 0, 0
	for j < n {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
