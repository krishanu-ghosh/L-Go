package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(listOfNumbers []int) int {
	max := listOfNumbers[0]
	for _, number := range listOfNumbers {
		if max < number {
			max = number
		}
	}
	return max
}

func min(listOfNumbers []int) int {
	min := listOfNumbers[0]
	for _, number := range listOfNumbers {
		if min > number {
			min = number
		}
	}
	return min
}

func main() {
	fmt.Print("Enter list of number : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	nums := strings.Fields(input)
	var numlist []int
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		numlist = append(numlist, n)
	}
	fmt.Println(numlist)
	fmt.Println("Max : ", max(numlist))
	fmt.Println("Min : ", min(numlist))
}
