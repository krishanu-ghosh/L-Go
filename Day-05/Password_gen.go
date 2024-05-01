package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func intInput(text string) int {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	i, _ := strconv.Atoi(input)
	return i
}

func shuffle(password string) string {
	//Fisher-Yates shuffle algorithm
	list := strings.Split(password, "")
	for i := len(list) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	password = strings.Join(list, "")
	return password
}

//func stringToSlice(input string) []string {
//	return strings.Split(input, "")
//}
//
//func sliceToString(input []string) string {
//	return strings.Join(input, "")
//}

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
		"v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q",
		"R",
		"S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	symbols := []string{"!", "#", "$", "%", "&", "(", ")", "*", "+"}
	nrLetters := intInput("Enter the number of letters you would like in your password ? : ")
	nrNumbers := intInput("How many numbers would you like? : ")
	nrSymbols := intInput("How many symbols would you like? : ")
	//fmt.Printf("letters: %d\nnumbers: %d\nsymbols: %d\n, ", nr_letters, nr_numbers, nr_symbols)
	var password string
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < nrLetters; i++ {
		password += letters[rand.Intn(len(letters))]
	}

	for i := 0; i < nrSymbols; i++ {
		password += symbols[rand.Intn(len(symbols))]
	}

	for i := 0; i < nrNumbers; i++ {
		password += numbers[rand.Intn(len(numbers))]
	}

	fmt.Println(password)
	fmt.Println(shuffle(password))
}
