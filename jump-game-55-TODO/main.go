package main

import "fmt"

/*

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

Initial: [4:true]
current_pos (3)+ value[current-pos](1)=> Is in map, if yes add it to map as true
[4: true, 3: true]

current_pos (2) + value[current_pos](1) => whatever is present in map, if yes add index to map as true
[4: true, 3: true, 2: true]

current_pos (1) + value[current_pos](3) => whatever is present in map, if yes add index to map as true
[4: true, 3: true, 2: true, 1: true]

current_pos (0) + value[current_pos](2) => whatever is present in map, if yes add index to map as true
[4: true, 3: true, 2: true, 1: true, 0: true]
*/

func canJump(nums []int) bool {
	/*
		jumpIndexMap := make(map[int]bool)
		lastIndex := len(nums) - 1 //1
		if nums[lastIndex] == 0 {
			return true
		}
		jumpIndexMap[lastIndex] = true

		for i := lastIndex - 1; i >= 0; i-- { //1-0
			searchIndex := i + nums[i] //1+0
			if jumpIndexMap[searchIndex] {
				jumpIndexMap[i] = true
				continue
			}
			jumpIndexMap[i] = false
		}
		return jumpIndexMap[0]
	*/

	i := 0
	lastIndex := len(nums) - 1
	for i < lastIndex {
		if i != lastIndex || i < lastIndex || nums[i] != 0 {
			i = i + nums[i]
		}
		return false
	}
}

func main() {
	nums := []int{2, 0, 2}

	fmt.Println(canJump(nums))
}
