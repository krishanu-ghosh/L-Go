package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dealCard() int {
	cards := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	rand.NewSource(time.Now().UnixNano())
	card := cards[rand.Intn(len(cards))]
	return card
}

func calculateScore(list []int) {

}

func main() {
	var userCards, computerCards []int
	for i := 1; i <= 2; i++ {
		userCards = append(userCards, dealCard())
		computerCards = append(computerCards, dealCard())
	}

	fmt.Println("User cards:", userCards)
	fmt.Println("Computer cards:", computerCards)
}
