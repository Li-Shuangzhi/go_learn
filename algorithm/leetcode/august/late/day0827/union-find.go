package main

//
//import "fmt"
//
//var parent []int
//
//func adapter() {
//	initFindUnion(9)
//	fmt.Println(isConnected(1, 5))
//	union(1, 5)
//	union(1,7)
//	fmt.Println(isConnected(1, 7))
//	fmt.Println(parent)
//}
//
//func union(x, y int) {
//	p := parent[x]
//	q := parent[y]
//	for i := 0; i < len(parent); i++ {
//		if parent[i] == p {
//			parent[i] = q
//		}
//	}
//}
//
//func isConnected(x, y int) bool {
//	return parent[x] == parent[y]
//}
//
//func initFindUnion(n int) {
//	parent = make([]int, n)
//	for i := 0; i < n; i++ {
//		parent[i] = i
//	}
//}
