package main

import (
	"fmt"
	"math"
)

// Write an algorithm to determine if a number is "happy".
// A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

// Example:

// Input: 19
// Output: true
// Explanation:
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1
func main() {
	result := isHappy(2)
	fmt.Println(result)
}

func isHappy(n int) bool {
	seenInts := make([]int, 0)
	return isHappyWithSum(n, seenInts)
}

func isHappyWithSum(n int, seenInts []int) bool {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(2)))
		fmt.Printf("%d^2 + ", digit)
		n /= 10
	}

	fmt.Printf("= %d\r\n", sum)
	if sum == 1 {
		return true
	}

	for _, v := range seenInts {
		if sum == v {
			return false
		}

	}
	seenInts = append(seenInts, sum)

	return isHappyWithSum(sum, seenInts)
}
