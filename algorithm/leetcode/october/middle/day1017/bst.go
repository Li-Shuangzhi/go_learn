package main

import (
	. "gocode/algorithm/utils"
)

func main() {

}
func kthSmallest(root *TreeNode, k int) int {
	vector := make([]int, 0)
	vector = help(root, vector)
	return vector[k-1]
}

func help(node *TreeNode, vector []int) []int {
	if node == nil {
		return vector
	}
	vector = help(node.Left, vector)
	vector = append(vector, node.Val)
	vector = help(node.Right, vector)
	return vector
}
