package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]float64, 0)
	p, q := 0, 0
	for p < len(nums1) && q < len(nums2) {
		n1 := nums1[p]
		n2 := nums2[q]
		if n1 <= n2 {
			nums = append(nums, float64(n1))
			p++
		} else {
			nums = append(nums, float64(n2))
			q++
		}
	}
	if p == len(nums1) {
		for i := q; i < len(nums2); i++ {
			nums = append(nums, float64(nums2[i]))
		}
	} else {
		for i := p; i < len(nums1); i++ {
			nums = append(nums, float64(nums1[i]))
		}
	}
	if len(nums)%2 == 0 {
		return (nums[len(nums)/2] + nums[len(nums)/2-1]) / 2
	} else {
		return nums[len(nums)/2]
	}
}
