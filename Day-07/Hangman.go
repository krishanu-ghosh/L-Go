package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func prepWord(word string) []string {
	word = strings.ToLower(strings.TrimSpace(word))
	return strings.Split(word, "")
}

func createBlanks(word string) []string {
	var blanks string
	for i := 1; i <= len(word); i++ {
		blanks += "_"
	}
	blanksList := strings.Split(blanks, "")
	return blanksList
}

func listToString(list []string) string {
	var theString string
	theString = strings.Join(list, "")
	return theString
}

//go:embed dictionary.txt
var filePath string

func main() {
	filePath = "dictionary.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	wordsByte, err1 := io.ReadAll(file)
	if err1 != nil {
		fmt.Println("Error: ", err1)
	}

	wordsList := strings.Split(string(wordsByte), "\n")
	//fmt.Println(wordsList)
	rand.NewSource(time.Now().UnixNano())
	word := wordsList[rand.Intn(len(wordsList))]
	word_L := prepWord(word)
	//fmt.Println(string(word[0]), word)
	fmt.Printf("Hint: The word starts with %v\n", string(word[0]))
	blanks := createBlanks(word)

	whileOn := true
	livesLeft := len(word_L) - 1
	fmt.Printf("You have %v live left.\n", livesLeft)
	fmt.Println(blanks)
	var guesses []string

	for whileOn && livesLeft > 0 {
		guessTag := false
		fmt.Printf("Your guesses so far : %v\n", guesses)
		fmt.Print("Guess letter : ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Error : ", err)
		}
		fmt.Println("\n\n")
		guesses = append(guesses, input)
		for i, v := range word_L {
			if v == input {
				blanks[i] = input
				guessTag = true
			}
		}
		if guessTag == false {
			livesLeft -= 1
			if livesLeft != 0 {
				fmt.Printf("Wrong guess !!! You have %v live left.\n", livesLeft)
			} else {
				whileOn = false
				fmt.Println("Game Over !!! You suck lol.")
			}
		} else {
			fmt.Printf("Correct guess !!! You have %v live left.\n", livesLeft)
		}
		if !strings.Contains(listToString(blanks), "_") {
			whileOn = false
		}
		if whileOn == false && livesLeft != 0 {
			fmt.Printf("You won !!! The word was %v \n", listToString(word_L))
			fmt.Println(blanks)
		} else if whileOn == false {
			fmt.Println("The word was", listToString(word_L))
		} else {
			fmt.Println(blanks)
		}
	}
}
