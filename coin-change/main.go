package main

import (
	"fmt"
	"sort"
)

// You are given coins of different denominations and a total amount of money amount.
// Write a function to compute the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.
func main() {
	// [186,419,83,408]
	// 6249
	coins := []int{1, 2, 5}
	resp := coinChange(coins, 11)
	fmt.Println(resp)

}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	result := 0
	leftOver := amount
	for i := len(coins) - 1; i >= 0; i-- {
		for j := 0; j < 100; j++ {
			d := leftOver - coins[i]

			if d == 0 {
				result++
				return result
			}

			if d < 0 {
				fmt.Println("d<0", d, leftOver, coins[i])
				break
			}

			leftOver = leftOver - coins[i]
			result++
			fmt.Printf("coin %d, leftover %d\n", coins[i], leftOver)
		}
	}

	return -1

}
