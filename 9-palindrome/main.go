package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	reverse, remainder := 0, 0
	originalNum := x
	for originalNum > 0 {
		remainder = originalNum % 10
		originalNum = originalNum / 10
		reverse = reverse*10 + remainder
	}
	fmt.Println(reverse)
	if x == reverse {
		return true
	}
	return false

}
