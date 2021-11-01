package day1101

import "testing"

func TestDistributeCandie(t *testing.T) {
	//candies := distributeCandies([]int{1, 1, 2, 2, 3, 3})
	//t.Log(candies)
	//candies = distributeCandies([]int{1,1,2,3})
	//t.Log(candies)
	//candies = distributeCandies([]int{0,0,14,0,10,0,0,0})
	//t.Log(candies)
	res := longestPalindromeSubseq("bbbab")
	t.Log(res)
	res = longestPalindromeSubseq("cbbd")
	t.Log(res)
	res = longestPalindromeSubseq("c")
	t.Log(res)
	res = longestPalindromeSubseq("cbbd")
	t.Log(res)
}
