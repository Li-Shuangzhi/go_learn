package day1114

//
//var ways []int
//var res [][]string
//var size int
//var str string
//
//func partition(s string) [][]string {
//	str = s
//	size = len(s)
//	res = make([][]string, 0)
//	backTrace(0)
//	return res
//}
//
//func backTrace(index int) {
//	if index == size-1 {
//		stringArray := getStringArray()
//		if check(stringArray) {
//			tmp := make([]int,len(ways))
//			copy(tmp, ways)
//			res = append(res, stringArray)
//		}
//		return
//	}
//	ways = append(ways, index)
//	backTrace(index + 1)
//	ways = ways[:len(ways)-1]
//	backTrace(index + 1)
//}
//
//func check(array []string) bool {
//	for i := 0; i < len(array); i++ {
//		if !isPlalindrome(array[i]) {
//			return false
//		}
//	}
//	return true
//}
//
//func isPlalindrome(s string) bool {
//	left, right := 0, len(s) -1
//	for left < right {
//		if s[left] == s[right] {
//			left++
//			right--
//		}else{
//			return false
//		}
//	}
//	return true
//}
//
//func getStringArray() []string {
//	result := make([]string, 0)
//	if len(ways) == 0 {
//		return []string{str}
//	}
//	t := 0
//	for i := 0; i < len(ways); i++ {
//		middle := ways[i] + 1
//		result = append(result, str[t:middle])
//		t = middle
//	}
//	result = append(result,str[t:])
//	return result
//}
