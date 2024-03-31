package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 7
	a[1] = 7
	a[2] = 7
	a[3] = 7
	a[4] = 7
	fmt.Println("--------------------------------")
	fmt.Println(a)
	fmt.Println(b)
}
