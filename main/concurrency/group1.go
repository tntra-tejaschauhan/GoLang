package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func Main1() {
	// the WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note : if a WaitGroup is explicitly passed into functions, it should be done by pointer.

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}
	// block until all the goroutines started by wg are done. A goroutine is done when the funcion it invokes returns.
	wg.Wait()
}
