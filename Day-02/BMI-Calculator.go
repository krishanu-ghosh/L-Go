package main

import (
	"fmt"
	"math"
)

func main() {
	var height float64
	var weight float64
	fmt.Print("Enter your height in metres : ")
	fmt.Scanf("%f", &height)
	fmt.Print("Enter your weight in kilograms : ")
	fmt.Scanf("%f", &weight)
	//fmt.Println(height, weight)

	bmi := weight / (math.Pow(height, 2))
	fmt.Println("Your BMI is ", bmi)

	if bmi <= 18.5 {
		fmt.Println("You are underweight")
	} else if 18.5 < bmi && bmi <= 25 {
		fmt.Println("You have Normal Weight")
	} else if 25 < bmi && bmi <= 30 {
		fmt.Println("You are Overweight.\nConsult a Dietitian")
	} else if bmi > 30 {
		fmt.Println("You are OBESE.\nDon't eat like a PIG you IDIOT.")
	}

}
