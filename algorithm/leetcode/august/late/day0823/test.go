package main

func main() {

}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	root = trimBSTLeft(root, low)
	return trimBSTRight(root, high)
}
func trimBSTLeft(root *TreeNode, low int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > low {
		root.Left = trimBSTLeft(root.Left, low)
		return root
	}
	if root.Val < low {
		root.Left = nil
		return trimBSTLeft(root.Right, low)
	}
	root.Left = nil
	return root
}

func trimBSTRight(root *TreeNode, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < high {
		root.Right = trimBSTRight(root.Right, high)
		return root
	}
	if root.Val > high {
		root.Right = nil
		return trimBSTRight(root.Left, high)
	}
	root.Right = nil
	return root
}
