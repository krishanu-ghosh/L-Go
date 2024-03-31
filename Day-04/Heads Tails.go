package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNum := rand.Intn(2)

	if randomNum == 0 {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}
}
