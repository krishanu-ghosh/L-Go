package main

import "fmt"

func main() {
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
	fmt.Println("====================================")
	fmt.Println("Inverted")

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
