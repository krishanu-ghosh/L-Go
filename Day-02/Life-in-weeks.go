package main

import "fmt"

func main() {
	fmt.Printf("Enter your age (in years) : ")
	var ageInput int
	fmt.Scanf("%d", &ageInput)
	yearsLeft := 60 - ageInput
	fmt.Printf("You have %v years, %v weeks and %v days left.\n", yearsLeft, yearsLeft*52, yearsLeft*365)
}
