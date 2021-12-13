package day1114

import "strconv"

var str string
var result []string
var ways []int

func restoreIpAddresses(s string) []string {
	for _, v := range s {
		if v > '9' || v < '0' {
			return nil
		}
	}
	str = s
	result = make([]string, 0)
	ways = make([]int, 0)
	backTrace(0)
	return result
}

func backTrace(index int) {
	if len(ways) == 3 {
		if isValid() {
			s := getIPStr()
			result = append(result, s)
		}
		return
	}
	if index > len(str)-2 {
		return
	}
	ways = append(ways, index)
	backTrace(index + 1)
	ways = ways[:len(ways)-1]
	backTrace(index + 1)
}

func getIPStr() string {
	bytes := make([]byte, 0)
	set := make(map[int]bool)
	for _, v := range ways {
		set[v] = true
	}
	for i := 0; i < len(str); i++ {
		bytes = append(bytes, str[i])
		if set[i] {
			bytes = append(bytes, '.')
		}
	}
	return string(bytes)
}

func isValid() bool {
	// 0,1,2 三个切割点
	t := 0
	for i := 0; i < len(ways); i++ {
		end := ways[i] + 1
		r := str[t:end]
		if !check(r) {
			return false
		}
		t = end
	}
	if !check(str[t:]) {
		return false
	}
	return true
}

func check(s string) bool {
	//长度大于三
	if len(s) > 3 {
		return false
	}
	//前导零
	i, _ := strconv.Atoi(s)
	ss := strconv.Itoa(i)
	if len(ss) < len(s) {
		return false
	}
	if i > 255 {
		return false
	}
	return true
}
