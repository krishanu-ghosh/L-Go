package main

import (
	"fmt"
	"math"
)

func minCoins(amount int, denominations []int) int {
	// Create a slice to store the minimum number of coins needed for each amount
	dp := make([]int, amount+1)
	// Initialize all values to infinity except for dp[0] which is 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	// Iterate through each amount from 1 to the given amount
	for i := 1; i <= amount; i++ {
		// Iterate through each denomination
		for _, coin := range denominations {
			// If the coin denomination is less than or equal to the current amount
			if coin <= i {
				// Update the minimum number of coins needed for the current amount
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example usage:
	amount := 12
	denominations := []int{1, 5, 10}
	fmt.Println("Minimum number of coins needed:", minCoins(amount, denominations))
}
