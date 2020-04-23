package main

import "fmt"

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:

// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
func main() {
	in1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans1 := maxSubArray(in1)
	if ans1 != 6 {
		fmt.Printf("In %v, wanted 6, got %d \r\b", in1, ans1)
	}

	in2 := []int{-2, -1}
	ans2 := maxSubArray(in2)
	if ans2 != -1 {
		fmt.Printf("In %v, wanted -1, got %d \r\n", in2, ans2)
	}

	in3 := []int{-1, -2}
	ans3 := maxSubArray(in3)
	if ans3 != -1 {
		fmt.Printf("In %v, wanted -1, got %d \r\n", in3, ans2)
	}

	in4 := []int{-1, 0, -2}
	ans4 := maxSubArray(in4)
	if ans4 != 0 {
		fmt.Printf("In %v, wanted 0, got %d \r\n", in4, ans4)
	}

}

func maxSubArray(nums []int) int {
	maxSoFar := -9999
	maxEndingHere := 0
	if len(nums) == 1 {
		return nums[0]
	}

	for _, v := range nums {
		maxEndingHere += v

		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}

	}
	return maxSoFar

}
