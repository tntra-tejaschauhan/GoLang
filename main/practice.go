// Go Philosophy in One File
// Read top to bottom. Every concept builds on the previous one.

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

//
// 1. BASIC TYPES, VARIABLES, CONSTANTS
//

const AppName = "Understand Go"

func basics() {
	// explicit type
	var a int = 10

	// inferred type
	b := 20

	// multiple assignment
	c, d := 30, "hello"

	fmt.Println(a, b, c, d)
}

//
// 2. FUNCTIONS (multiple return values are normal in Go)
//

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

//
// 3. STRUCTS + METHODS
//

type User struct {
	Name string
	Age  int
}

// method with value receiver
func (u User) Greet() string {
	return "Hello, my name is " + u.Name
}

// method with pointer receiver (can modify struct)
func (u *User) Birthday() {
	u.Age++
}

//
// 4. INTERFACES (implicit implementation)
//

type Greeter interface {
	Greet() string
}

func greetSomeone(g Greeter) {
	fmt.Println(g.Greet())
}

//
// 5. SLICES & MAPS
//

func collections() {
	// slice (dynamic array)
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)

	// map (hash table)
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	fmt.Println(numbers, ages["Alice"])
}

//
// 6. GENERICS (Go 1.18+)
//

func Sum[T int | float64](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

//
// 7. CONCURRENCY: GOROUTINES + CHANNELS
//

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Millisecond * 200)
		results <- job * 2
	}
}

func prec_concurrency() {
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	// start goroutines
	for i := 1; i <= 2; i++ {
		go worker(i, jobs, results)
	}

	// send jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)

	// receive results
	for i := 1; i <= 3; i++ {
		fmt.Println("result:", <-results)
	}
}

//
// 8. DEFER + SYNC (real-world Go pattern)
//

func safeCounter() {
	var mu sync.Mutex
	counter := 0

	increment := func() {
		mu.Lock()
		defer mu.Unlock() // always runs, even on panic
		counter++
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}

//
// 9. MAIN = PROGRAM ENTRY POINT
//

func main1() {
	fmt.Println(AppName)

	basics()

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("divide result:", result)
	}

	user := User{Name: "Alice", Age: 30}
	greetSomeone(user)

	user.Birthday()
	fmt.Println("age after birthday:", user.Age)

	collections()

	fmt.Println("sum ints:", Sum([]int{1, 2, 3}))
	fmt.Println("sum floats:", Sum([]float64{1.5, 2.5}))

	prec_concurrency()

	safeCounter()
}
