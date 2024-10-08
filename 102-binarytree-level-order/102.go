package binarytreelevelorder102

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	levelMap := make(map[int][]int)

	traverse(levelMap, 0, root)

	returnResponse := make([][]int, len(levelMap))

	for key, value := range levelMap {
		returnResponse[key] = value
	}

	return returnResponse
}

func traverse(levelMap map[int][]int, level int, root *TreeNode) {

	if root == nil {
		return
	}
	levelMap[level] = append(levelMap[level], root.Val)

	traverse(levelMap, level+1, root.Left)
	traverse(levelMap, level+1, root.Right)
}
