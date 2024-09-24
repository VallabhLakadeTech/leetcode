package main

import (
	"math"
)

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

Constraints:

-231 <= x <= 231 - 1
*/

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	reverse := 0
	for x > 0 {
		remainder := x % 10
		x = x / 10
		reverse = reverse*10 + remainder
	}

	reverse = reverse * sign

	if reverse < math.MinInt32 || reverse > math.MaxInt32 {
		return 0
	}
	return reverse

}
