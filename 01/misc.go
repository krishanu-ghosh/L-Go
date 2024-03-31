package main

import "fmt"

func main() {
	// a := [3]int{4, 5, 6}
	// b := []int{7, 8, 9}
	// c := make([]string, 10)
	// fmt.Println(c)
	// fmt.Printf("Value of a is %v and type is %T\n", a, a)
	// fmt.Printf("Value of a is %v and type is %T\n", b, b)
	// a := 45
	a := []int{1, 2, 3, 4, 5}
	// b := a
	b := make([]int, 5)
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("-----------------------")

	// a = 46
	// a = []int{6, 7, 8, 9, 0}
	a[0] = 3
	fmt.Println(a)
	fmt.Println(b)
}
