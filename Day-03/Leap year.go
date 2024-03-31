package main

import "fmt"

func main() {
	var yearToCheck int
	fmt.Printf("Enter yaer to check : ")
	fmt.Scanln(&yearToCheck)

	if yearToCheck%4 == 0 {
		if yearToCheck%100 == 0 {
			if yearToCheck%400 == 0 {
				fmt.Printf("The year %v is a leap year.\n", yearToCheck)
			} else {
				fmt.Printf("The year %v is not a leap year.\n", yearToCheck)
			}
		} else {
			fmt.Printf("The year %v is a leap year.\n", yearToCheck)
		}
	} else {
		fmt.Printf("The year %v is not a leap year.\n", yearToCheck)
	}
}
