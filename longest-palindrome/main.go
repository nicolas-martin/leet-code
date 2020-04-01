package main

import "fmt"

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

// Example 1:
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
// Example 2:
// Input: "cbbd"
// Output: "bb"

func main() {
	resp := longestPalindrome("babad")
	fmt.Println(resp)

}

func longestPalindrome(s string) string {

	return ""
}

func isPalindrome(s string) bool {
	arr := make([int]int, 128)
	for _, c := range s {
		if arr[c] == 0 {
			arr[c]++
		}
	}

	return false
}
