package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	help(root)
	return result
}
func help(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	if left+right > result {
		result = left + right
	}
	return max(left, right) + 1
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
