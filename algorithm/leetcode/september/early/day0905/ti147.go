package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) (res *ListNode) {
	if head == nil {
		return nil
	}
	p, q := res, head
	pre := res
	for q != nil {
		for p != nil {
			if q.Val >= p.Val {
				p = p.Next
				pre = p
			} else {
				pre.Next = &ListNode{q.Val, p}
			}
		}
	}

	return nil
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	for k > 0 {
		p = p.Next
		if p == nil {
			return nil
		}
		k--
	}
	cur := head
	for p != nil {
		p = p.Next
		cur = cur.Next
	}
	return cur
}

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return letters[left]
}
