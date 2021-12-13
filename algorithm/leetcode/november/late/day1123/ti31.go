package day1123

func nextPermutation(nums []int) {
	size := len(nums)
	left := -1
	for i := size - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			left = i
			break
		}
	}
	if left >= 0 {
		for i := size - 1; i >= 0; i-- {
			if nums[i] > nums[left] {
				nums[i], nums[left] = nums[left], nums[i]
				break
			}
		}
	}
	reverse(nums[left+1:])
}

func reverse(a []int) {
	left, right := 0, len(a)-1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}
