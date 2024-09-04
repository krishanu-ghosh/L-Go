package main

import (
	"io"
	"net/http"
	"sync"
)

func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Discard the data (if you just want to generate traffic)
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "http://speedtest.tele2.net/100MB.zip" // Example file, make sure it's safe to download repeatedly.

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ { // Adjust the number of goroutines based on your needs.
		wg.Add(1)
		go downloadFile(url, &wg)
	}

	wg.Wait()
}
