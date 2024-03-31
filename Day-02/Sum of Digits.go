package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	fmt.Printf("Enter a number : ")
	fmt.Scan(&num)
	//fmt.Printf("\nThe sum of %v and %v is %v", int(num[0]), int(num[1]), int(num[0])+int(num[1]))
	firstDigit, _ := strconv.Atoi(string(num[0]))
	secondDigit, _ := strconv.Atoi(string(num[1]))
	fmt.Printf("\nThe sum of %v and %v is %v\n", firstDigit, secondDigit, firstDigit+secondDigit)
}
