package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 6, 6, 2}
	resp := topKFrequent(nums, 2)
	fmt.Println(resp)

}
func topKFrequent(nums []int, k int) []int {
	ints := make(map[int][2]int)
	for i, v := range nums {
		if ints[i][1] == 0 {
			ints[i][1] = v
			continue
		}
		ints[i] = ints[i] + 1
	}
	result := []int{}

	for i, v := range ints {
		if v >= k {
			result = append(result, i)
		}
	}

	return result

}
