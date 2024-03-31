package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter number to check : ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Printf("The number %v is even\n", num)
	} else {
		fmt.Printf("The number %v is odd\n", num)
	}
}
