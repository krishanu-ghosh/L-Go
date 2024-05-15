package main

import "fmt"

func main() {
	student_scores := map[string]int{
		"Tom":   81,
		"Dick":  96,
		"Harry": 75,
		"Joe":   66,
		"Jane":  33,
	}

	for name, score := range student_scores {
		if score > 90 {
			fmt.Printf("%v : Outstanding\n", name)
		} else if score > 80 {
			fmt.Printf("%v : Exceeds expectation\n", name)
		} else if score > 70 {
			fmt.Printf("%v : Above average\n", name)
		} else if score > 40 {
			fmt.Printf("%v : Passed\n", name)
		} else {
			fmt.Printf("%v : Failed\n", name)
		}
	}
}
