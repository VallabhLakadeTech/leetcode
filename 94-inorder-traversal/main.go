package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	inorderSlice := []int{}
	if root != nil {
		inorderSlice = append(inorderSlice, inorderTraversal(root.Left)...)
		inorderSlice = append(inorderSlice, root.Val)
		inorderSlice = append(inorderSlice, inorderTraversal(root.Right)...)
	}
	return inorderSlice
}
