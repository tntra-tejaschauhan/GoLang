package main

import (
	"fmt"
	"time"
)

func worker(dataChan <-chan int, stopChan <-chan struct{}) {
	for {
		select {
		case job := <-dataChan:
			// Process the job
			fmt.Printf("Processing job: %d\n", job)
		case <-stopChan:
			// Received stop signal, exit the goroutine
			fmt.Println("Worker received stop signal, exiting.")
			return
		}
	}
}

func main1() {
	dataChan := make(chan int)
	stopChan := make(chan struct{}) // The signal channel

	go worker(dataChan, stopChan)

	// Send some jobs
	for i := 1; i <= 5; i++ {
		dataChan <- i
	}

	// Give time for some processing
	time.Sleep(500 * time.Millisecond)

	// Signal the worker to stop by closing the stop channel
	close(stopChan)

	// Wait a moment to see the exit message
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main function exiting.")
}
