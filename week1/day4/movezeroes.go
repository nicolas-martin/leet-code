package main

import "fmt"

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Example:

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Note:

// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.
func main() {
	in := []int{0, 1, 0, 3, 12}
	// out := []int{1, 3, 12, 0, 0}
	moveZeroes(in)
	fmt.Printf("%v\v\n", in)

}
func moveZeroes(nums []int) {

	count := 0
	for k := range nums {
		if nums[k] != 0 {
			nums[count] = nums[k]
			count++
		}
	}

	for count != len(nums) {
		nums[count] = 0
		count++
	}

}
