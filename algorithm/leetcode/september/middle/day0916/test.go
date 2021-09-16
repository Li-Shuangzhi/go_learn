package main

import . "gocode/algorithm/utils"

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	p := preorder[0]
	node := &TreeNode{Val: p}
	t := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == p {
			t = i
		}
	}
	node.Left = buildTree(preorder[1:1+t], inorder[:t])
	node.Right = buildTree(preorder[t+1:], inorder[t+1:])
	return node
}
