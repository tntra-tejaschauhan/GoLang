package mutex

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()         // 🔒 lock
	counter = counter + 1
	mu.Unlock()       // 🔓 unlock
}

func main2() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
