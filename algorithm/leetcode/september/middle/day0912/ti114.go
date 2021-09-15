package main

import (
	. "gocode/algorithm/utils"
)

func main() {

}

var last *TreeNode = nil

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Right = last
	root.Left = nil
	last = root
}
