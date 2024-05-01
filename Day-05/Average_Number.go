package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func averageNumber(numlist []int) float64 {
	total := 0
	for _, n := range numlist {
		total = total + n
	}
	var avg float64
	avg = float64(total) / float64(len(numlist))
	return avg
}

func main() {
	//fmt.Printf("Enter numbers separated by space : ")
	// Prompt the user to enter numbers
	fmt.Print("Enter nums: ")

	// Read input from standard input (stdin)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove trailing newline character
	input = strings.TrimSpace(input)

	// Split the input string by whitespace
	nums := strings.Fields(input)
	var numList []int

	// Print the numbers entered by the user
	//fmt.Println("Entered nums:", nums)
	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		numList = append(numList, num)
	}

	fmt.Println(numList)
	fmt.Println(averageNumber(numList))

}
