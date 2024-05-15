package main

import (
	"fmt"
	"os"
)

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func sub(n1, n2 float64) float64 {
	return n1 - n2
}

func mul(n1, n2 float64) float64 {
	return n1 * n2
}

func div(n1, n2 float64) float64 {
	if n2 == 0 {
		fmt.Println("Error: Division by zero")
		os.Exit(1)
	}
	return n1 / n2
}

func main() {
	operations := map[string]func(float64, float64) float64{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	var n1, n2 float64
	fmt.Printf("Enter first number: ")
	fmt.Scanf("%f", &n1)
	var ops string
	fmt.Printf("Enter operator (  +  -  *  /  ) : ")
	fmt.Scanf("%s", &ops)
	fmt.Printf("Enter second number: ")
	fmt.Scanf("%f", &n2)
	result := operations[ops](n1, n2)
	fmt.Printf("Result: %.2f\n", result)

}
