package main

import "fmt"

func main() {
	var bill, tip, people float64
	fmt.Printf("Enter the bill amount : ")
	fmt.Scanln(&bill)
	fmt.Printf("Enter the tip percentage : ")
	fmt.Scanln(&tip)
	fmt.Printf("Enter the number of people to split the bill : ")
	fmt.Scanln(&people)

	totalBill := bill + ((tip / 100) * bill)
	split := totalBill / people

	fmt.Printf("Each person have to pay %v\n", split)

}
