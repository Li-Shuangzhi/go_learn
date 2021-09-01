package main

func decode(encoded []int) []int {
	n := len(encoded)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	t := 0
	for i := 1; i < n; i += 2 {
		t ^= encoded[i]
	}
	result := make([]int, n+1)
	result[0] = total ^ t
	for i := 1; i <= n; i++ {
		result[i] = encoded[i-1] ^ result[i-1]
	}
	return result
}
