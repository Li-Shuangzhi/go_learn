package main

import "fmt"

func main() {
	roman := intToRoman(1994)
	fmt.Println(roman)
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	res := ""
	for i := 0; i < len(valueSymbols); i++ {
		t := num / valueSymbols[i].value
		for j := 0; j < t; j++ {
			res += valueSymbols[i].symbol
		}
		num %= valueSymbols[i].value
	}
	return res
}
