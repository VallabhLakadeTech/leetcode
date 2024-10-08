package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	postorderSlice := []int{}
	if root != nil {
		postorderSlice = append(postorderSlice, postorderTraversal(root.Left)...)
		postorderSlice = append(postorderSlice, postorderTraversal(root.Right)...)
		postorderSlice = append(postorderSlice, root.Val)
	}
	return postorderSlice
}
