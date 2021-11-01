package main

import "fmt"

func main() {
	connection := findRedundantConnection([][]int{
		{1, 4},
		{3, 4},
		{1, 3},
		{1, 2},
		{4, 5},
	})
	fmt.Println(connection)
}

var ids []int
var ranks []int

func findRedundantConnection(edges [][]int) []int {
	initIds(len(edges))
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		if find(x) == find(y) {
			return edges[i]
		} else {
			union(x, y)
		}
	}
	return nil
}

func initIds(size int) {
	ids = make([]int, size+1)
	ranks = make([]int, size+1)
	for i := 1; i < len(ids); i++ {
		ids[i] = i
		ranks[i] = 1
	}
}

func union(x, y int) {
	i := find(x)
	j := find(y)
	if ranks[i] > ranks[j] {
		ids[j] = i
		ranks[i] += ranks[j]
	} else {
		ids[i] = j
		ranks[j] += ranks[i]
	}
}

func find(x int) int {
	if ids[x] == x {
		return x
	}
	return find(ids[x])
}
