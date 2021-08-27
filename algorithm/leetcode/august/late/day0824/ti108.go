package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	return help(nums, 0, len(nums)-1)
}

func help(nums []int, left, right int) *TreeNode {
	mid := left + (right-left)/2
	if right-left < 0 {
		return nil
	}
	if right-left == 0 {
		return &TreeNode{Val: nums[mid]}
	} else {
		left := help(nums, left, mid-1)
		right := help(nums, mid+1, right)
		return &TreeNode{Val: nums[mid], Left: left, Right: right}
	}
}
