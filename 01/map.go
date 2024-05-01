package main

import "fmt"

func main() {
	//m := make(map[string]int)
	//n := map[string]int{"one": 1, "two": 2, "three": 3}
	o := make(map[string][]int)
	o["one"] = []int{1, 2, 3}
	fmt.Println(o)
}
