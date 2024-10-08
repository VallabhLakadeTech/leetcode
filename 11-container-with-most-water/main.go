package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {

	maxArea := 0
	for start, end := 0, len(height)-1; start < end; {
		area := 0
		if height[start] < height[end] {
			area = height[start] * (end - start)
			start++
		} else if height[start] >= height[end] {
			area = height[end] * (end - start)
			end--
		}
		maxArea = max(maxArea, area)
	}
	return maxArea

}
