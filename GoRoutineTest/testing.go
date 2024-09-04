package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINE\t", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(1)
	go core1(&wg)
	core2()
	wg.Wait()
}

func core1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("core 1 :", i)
	}
}

func core2() {
	for i := 0; i < 10; i++ {
		fmt.Println("core 2 :", i)
	}
}
