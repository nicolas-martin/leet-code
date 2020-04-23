package main

import "fmt"

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

// Example 1:

// Input: [2,2,1]
// Output: 1
// Example 2:

// Input: [4,1,2,1,2]
// Output: 4

func main() {
	e1 := []int{2, 2, 1}
	single := singlenumber(e1)
	fmt.Println(single)

	e2 := []int{4, 1, 2, 1, 2}
	single2 := singlenumber(e2)
	fmt.Println(single2)
}

func singlenumber(arr []int) int {
	count := map[int]int{}

	for _, v := range arr {
		if _, ok := count[v]; !ok {
			count[v] = 1
		} else {
			count[v] = count[v] + 1
		}

	}

	for k, v := range count {
		if v == 1 {
			return k
		}

	}
	return 0

}
