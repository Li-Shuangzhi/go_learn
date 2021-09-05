package main

import "math/rand"

func main() {

}

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}

func rand7() int {
	return 1 + rand.Int()%7
}
