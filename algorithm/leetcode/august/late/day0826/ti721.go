package main

import (
	"fmt"
	"sort"
)

func main() {
	result := accountsMerge([][]string{
		{"Lily", "Lily4@m.co", "Lily5@m.co"},
		{"Lily", "Lily8@m.co", "Lily9@m.co"},
		{"Lily", "Lily15@m.co", "Lily16@m.co"},
		{"Lily", "Lily19@m.co", "Lily20@m.co"},
		{"Lily", "Lily6@m.co", "Lily7@m.co"},
		{"Lily", "Lily10@m.co", "Lily11@m.co"},
		{"Lily", "Lily5@m.co", "Lily6@m.co"},
		{"Lily", "Lily13@m.co", "Lily14@m.co"},
		{"Lily", "Lily9@m.co", "Lily10@m.co"},
		{"Lily", "Lily1@m.co", "Lily2@m.co"},
		{"Lily", "Lily3@m.co", "Lily4@m.co"},
		{"Lily", "Lily2@m.co", "Lily3@m.co"},
		{"Lily", "Lily11@m.co", "Lily12@m.co"},
		{"Lily", "Lily7@m.co", "Lily8@m.co"},
		{"Lily", "Lily12@m.co", "Lily13@m.co"},
		{"Lily", "Lily18@m.co", "Lily19@m.co"},
		{"Lily", "Lily17@m.co", "Lily18@m.co"},
		{"Lily", "Lily16@m.co", "Lily17@m.co"},
		{"Lily", "Lily14@m.co", "Lily15@m.co"},
		{"Lily", "Lily0@m.co", "Lily1@m.co"}})
	fmt.Println(result)
}

var numbers []int

func accountsMerge(accounts [][]string) (result [][]string) {
	numbers = make([]int, len(accounts))
	emailToNumber := make(map[string]int)
	numberToEmails := make(map[int][]string)
	result = make([][]string, len(accounts))

	for i := 0; i < len(accounts); i++ {
		numbers[i] = i
		numberToEmails[i] = make([]string, 0)
		result[i] = append(result[i], accounts[i][0])
	}

	for i := 0; i < len(accounts); i++ {
		account := accounts[i]
		for j := 1; j < len(account); j++ {
			emailToNumber[accounts[i][j]] = i
			numberToEmails[i] = append(numberToEmails[i], accounts[i][j])
		}
	}
	for i := 0; i < len(accounts); i++ {
		account := accounts[i]
		for j := 1; j < len(account); j++ {
			if value, ok := emailToNumber[account[j]]; ok {
				union(value, i)
			}
		}
	}

	for i := 0; i < len(numbers); i++ {
		mp := make(map[string]bool)
		tmp := make([]string, 0)
		for j := 0; j < len(numbers); j++ {
			if i == getRoot(j) {
				for k := 0; k < len(numberToEmails[j]); k++ {
					if has := mp[numberToEmails[j][k]]; !has {
						tmp = append(tmp, numberToEmails[j][k])
						mp[numberToEmails[j][k]] = true
					}
				}
			}
		}
		sort.Strings(tmp)
		result[i] = append(result[i], tmp...)
	}
	for j := 0; j < len(result); {
		if len(result[j]) == 1 {
			result = append(result[:j], result[j+1:]...)
		} else {
			j++
		}
	}
	fmt.Println(numbers)
	return
}
func getRoot(x int) int {
	if numbers[x] == x {
		return numbers[x]
	}
	numbers[x] = getRoot(numbers[x])
	return numbers[x]
}
func union(x, y int) {
	numbers[getRoot(x)] = getRoot(y)
}
