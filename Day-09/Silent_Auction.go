package main

import (
	"fmt"
	"strings"
)

func main() {
	addName := true
	people := map[string]int{}

	for addName {
		var name string
		var bid int
		fmt.Printf("Enter your name: ")
		fmt.Scanf("%s", &name)
		fmt.Printf("Enter your bid : ")
		fmt.Scanf("%d", &bid)
		people[name] = bid
		var flagToAdd string
		fmt.Printf("Enter another entry ? (Yes/No) : ")
		fmt.Scanf("%s", &flagToAdd)
		flagToAdd = strings.ToLower(flagToAdd)
		if flagToAdd == "no" || flagToAdd == "n" {
			addName = false
		} else {
			for i := 0; i < 20; i++ {
				fmt.Println()
			}
		}
	}

	var maxBid int
	var winner string
	for name, bid := range people {
		if bid > maxBid {
			maxBid = bid
			winner = name
		}
	}

	fmt.Printf("The winner is %s with a winning bid of $%v\n", winner, maxBid)

}
