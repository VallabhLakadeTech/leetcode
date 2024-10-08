package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	inorderSlice := []int{}
	if root != nil {
		inorderSlice = append(inorderSlice, root.Val)
		inorderSlice = append(inorderSlice, preorderTraversal(root.Left)...)
		inorderSlice = append(inorderSlice, preorderTraversal(root.Right)...)
	}
	return inorderSlice
}
