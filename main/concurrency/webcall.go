package concurrency

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchStatus(wg *sync.WaitGroup, url string) {
	// Tell WaitGroup this goroutine is done when function exits
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "-> error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(url, "-> status:", resp.StatusCode)
}

func main3() {
	websites := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://example.com",
	}

	var wg sync.WaitGroup

	for _, site := range websites {
		wg.Add(1)                 // increment counter
		go fetchStatus(&wg, site) // start goroutine
	}

	wg.Wait() // block until counter becomes 0
	fmt.Println("All requests completed")
}
