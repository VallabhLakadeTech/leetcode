package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNodeDepth := maxDepth(root.Left)
	rightNodeDepth := maxDepth(root.Right)
	return 1 + max(leftNodeDepth, rightNodeDepth)
}
