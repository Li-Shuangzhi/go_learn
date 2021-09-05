package main

func main() {

}

func findDuplicate1(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	tmp := 0
	for i := 1; i < len(nums); i++ {
		tmp += i
	}
	return sum - tmp
}

func findDuplicate(nums []int) int {
	mp := make(map[int]bool)
	for _, v := range nums {
		if mp[v] {
			return v
		} else {
			mp[v] = true
		}
	}
	return -1
}
