package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the input string
	input := `
		function myFunction() public view onlyOwner {
			// Function body
		}

		function anotherFunction() external onlyAdmin {
			// Function body
		}

		function payableFunction() public payable {
			// Function body
		}
	`

	// Define the regular expression pattern
	pattern := `\s*function.*?\{`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)

	// Print all matches
	for _, match := range matches {
		fmt.Println(match)
	}
}
