package main

import (
	"fmt"
	"strings"
)

func main() {
	var numInput string
	var temp int
	temp = 0

	fmt.Scanf("%s", &numInput)

	numList := strings.SplitAfter(numInput, "")

	fmt.Println(numList)
	fmt.Println(len(numList))

	for num := range numList {
		temp = temp + int(num)
		fmt.Println(temp)
	}
	fmt.Println(temp)
	fmt.Println("Average of the numbers = ", float32(temp)/float32(len(numList)))

}
