package main

func main() {

}

func searchRange(nums []int, target int) (res []int) {
	if len(nums) == 0 {
		res = append(res, -1, -1)
		return
	}
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[right] != target {
		res = append(res, -1)
	} else if nums[left] != nums[right] {
		res = append(res, right)
	} else {
		res = append(res, left)
	}
	left, right = 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		}
	}

	if nums[left] != target {
		res = append(res, -1)
	} else if nums[left] != nums[right] {
		res = append(res, left)
	} else {
		res = append(res, right)
	}
	return
}
