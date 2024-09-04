package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Killer() {
	var Tree [][]int
	elementsPerSlice := 10
	totalElements := 1_000_000_000_000

	for i := 0; i < totalElements; i += elementsPerSlice {
		var innerSlice []int
		for j := 0; j < elementsPerSlice; j++ {
			innerSlice = append(innerSlice, i+j+1)
		}
		Tree = append(Tree, innerSlice)
	}
	for _, v := range Tree {
		fmt.Println(v)
	}
	fmt.Println("\n====================================")
	fmt.Println("Inverted\n====================================\n")

	//Invert
	var invertedTree [][]int
	for i := len(Tree) - 1; i >= 0; i-- {
		var innerTree []int
		for j := len(Tree[i]) - 1; j >= 0; j-- {
			//fmt.Println(Tree[i][j])
			innerTree = append(innerTree, Tree[i][j])
		}
		invertedTree = append(invertedTree, innerTree)
	}
	for _, v := range invertedTree {
		fmt.Println(v)
	}

	query := 697

	for index, slice := range Tree {
		for i, v := range slice {
			if v == query {
				fmt.Printf("Row = %v \nColumn = %v \nQuery = %v \n", index+1, i+1, v)
				break
			}
		}
	}
}

func main() {
	url := "https://trollface.dk/"
	osName := runtime.GOOS
	if strings.Contains(osName, "linux") {
		cmd := exec.Command("xdg-open", url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error opening website:", err)
		}
	} else if strings.Contains(osName, "windows") {
		cmd := exec.Command("cmd", "/c", "start", url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error opening website:", err)
		}
	} else {
		cmd := exec.Command("open", url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error opening website:", err)
		}
	}
	time.Sleep(5 * time.Second)
	Killer()
}
