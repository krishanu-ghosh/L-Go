package main

import "fmt"

func main() {
	map1 := map[string]int{}
	map2 := map[string]int{}

	map1["key1"] = 5
	map2["key1"] = 3

	fmt.Println(map1["key1"])
	fmt.Println(map2["key1"])
}
