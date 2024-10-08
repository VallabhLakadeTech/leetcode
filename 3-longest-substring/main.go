package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	//abba
	//
	//abcde
	// aab
}

func lengthOfLongestSubstring(s string) int {
	inputString := s
	longestSubstring := 0
	start := 0
	charSet := make(map[rune]int)
	duplicateFound := false
	for index, char := range inputString {
		if i, ok := charSet[char]; !ok {
			charSet[char] = index
		} else {
			duplicateFound = true
			difference := index - start
			longestSubstring = max(difference, longestSubstring)
			if start < i+1 {
				start = i + 1
			}
			charSet[char] = index
		}
	}
	if !duplicateFound {
		return max(longestSubstring, len(charSet))
	}
	return max(longestSubstring, len(s)-start)

}
