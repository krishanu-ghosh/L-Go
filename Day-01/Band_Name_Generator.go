package main

import "fmt"

func main() {
	fmt.Println("----- Band Name Generator -----")
	var city, pet string
	fmt.Println("Enter name of your city : ")
	fmt.Scanf("%v", &city)
	fmt.Println("Enter name of your pet : ")
	fmt.Scanf("%v", &pet)
	fmt.Printf("Your band name could be %v %v\n", city, pet)

}
