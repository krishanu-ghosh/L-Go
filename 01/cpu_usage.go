package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"
	"sync"
)

func Killer2() {
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

func generatePrime(bits int) *big.Int {
	// Generate a prime number with the specified number of bits
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return p
}

func main() {
	//numCPU := 16 // Number of CPU cores
	numCPU := runtime.NumCPU()
	//fmt.Printf("numCPU: %d\n", numCPU)
	runtime.GOMAXPROCS(numCPU) // Set GOMAXPROCS to utilize all cores

	var wg sync.WaitGroup
	wg.Add(numCPU)

	for i := 0; i < numCPU; i++ {
		go func() {
			// Infinite loop to keep CPU busy
			for {
				// Generate large prime numbers
				prime1 := generatePrime(1024) // Generate a prime number with around 512 bits
				prime2 := generatePrime(1024) // Generate a prime number with around 1024 bits

				fmt.Println("Prime 1:", prime1)
				fmt.Println("Prime 2:", prime2)

				// Multiply the prime numbers
				result := new(big.Int).Mul(prime1, prime2)
				fmt.Println("Result of multiplication:", result)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
