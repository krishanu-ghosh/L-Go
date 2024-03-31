package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	var n int
	fmt.Printf("Enter number of people present : ")
	fmt.Scanln(&n)

	people := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Enter name #%d : ", i)
		var name string
		fmt.Scanln(&name)
		people = append(people, name)
	}
	//fmt.Println(len(people))
	fmt.Printf("%v will pay the bill.\n", people[rand.Intn(len(people))])
}
