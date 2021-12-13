package day1116

//
//var result [][]int
//var path []int
//var l int
//
//func combinationSum3(k int, n int) [][]int {
//	result = make([][]int, 0)
//	path = make([]int, 0)
//	l = k
//	backTrace(1,n,0)
//	return result
//}
//
//func min(x, y int) int {
//	if x >  y {
//		return y
//	}
//	return x
//}
//
//func backTrace(startIndex int, n int, sum int) {
//	if sum > n {
//		return
//	}
//	if sum == n && len(path) == l {
//		tmp := make([]int, len(path))
//		copy(tmp, path)
//		result = append(result, tmp)
//	}
//	for i := startIndex; i <= n; i++ {
//		if i > 9 {
//			return
//		}
//		path = append(path, i)
//		sum += i
//		backTrace(i+1, n, sum)
//		sum -= i
//		path = path[:len(path)-1]
//	}
//}
