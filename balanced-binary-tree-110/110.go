package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	leftTreeHeight := findHeight(root.Left)
	rightTreeHeight := findHeight(root.Right)

	diff := leftTreeHeight - rightTreeHeight
	if (diff == 0 || diff == -1 || diff == 1) && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func findHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftTreeHeight := 1 + findHeight(root.Left)
	rightTreeHeight := 1 + findHeight(root.Right)
	return max(leftTreeHeight, rightTreeHeight)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
