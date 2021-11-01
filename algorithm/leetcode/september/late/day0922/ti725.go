package main

import (
	"fmt"
	. "gocode/algorithm/utils"
)

func main() {

}

func splitListToParts(head *ListNode, k int) (res []*ListNode) {
	length := getLength(head)
	fmt.Println(length)
	//每组n个基本数
	n := length / k
	//m个余数
	m := length % k
	q := head
	for i := 0; i < k; i++ {
		if i+1 <= m {
			res = append(res, q)
			p := q
			for i := 0; i < n; i++ {
				p = p.Next
			}
			q = p.Next
			p.Next = nil
		} else {
			res = append(res, q)
			if q == nil {
				continue
			}
			p := q
			for i := 0; i < n-1; i++ {
				p = p.Next
			}
			q = p.Next
			p.Next = nil
		}
	}
	return
}

func getLength(head *ListNode) (res int) {
	p := head
	for p != nil {
		res++
		p = p.Next
	}
	return
}
