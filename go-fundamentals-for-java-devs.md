# Go Fundamentals for Java Developers

A quick reference guide for Java developers learning Go.

---

## Why Choose Go?

1. **Fast Build Time** - Compiles much faster than Java
2. **Fast Startup** - No JVM warmup needed
3. **Performance & Efficiency** - Low memory footprint, efficient runtime
4. **Concurrency Model** - Goroutines are lightweight (not OS threads)
5. **Static Typing & Compilation** - Type safety with compiled binaries

---

## Package Management

### Java vs Go

**Java:**
```java
// Package name and import both use full path
package com.example.myapp;
import com.example.utils.Helper;
```

**Go:**
```go
// Package name is local (last part), import path is full
package myapp

import "github.com/username/project/utils"
```

### Key Differences
- **Package name**: Just the folder name (e.g., `package main`, `package utils`)
- **Import path**: Full path to the package
- **Exported names**: Start with uppercase (e.g., `PublicFunc` vs `privateFunc`)

---

## Basic Syntax

### Variables

```go
// Explicit type
var name string = "John"
var age int = 25

// Type inference
var city = "New York"

// Short declaration (only inside functions)
count := 10
isActive := true

// Multiple declarations
var (
    firstName string = "Jane"
    lastName  string = "Doe"
    score     int    = 95
)
```

### Functions

```go
// Basic function
func add(a int, b int) int {
    return a + b
}

// Same type parameters (shorthand)
func multiply(x, y int) int {
    return x * y
}

// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return // returns sum and product automatically
}
```

---

## Error Handling (Different from Java!)

**Java uses exceptions:**
```java
try {
    result = divide(10, 0);
} catch (Exception e) {
    System.out.println(e.getMessage());
}
```

**Go uses explicit error returns:**
```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

### Best Practice
```go
// Check errors immediately
file, err := os.Open("data.txt")
if err != nil {
    log.Fatal(err) // Handle error right away
}
defer file.Close() // Ensure cleanup
```

---

## Pointers

Go has pointers like C/C++, unlike Java which hides them.

### Basics

```go
// Create a variable
x := 10

// Create a pointer to x
ptr := &x  // & gets the address

// Dereference pointer (get value)
value := *ptr  // * reads the value at address

fmt.Println(x)     // 10
fmt.Println(ptr)   // memory address like 0xc000012028
fmt.Println(*ptr)  // 10
```

### Why Use Pointers?

**Pass by value (copy):**
```go
func updateValue(num int) {
    num = 100 // Only changes the copy
}

x := 10
updateValue(x)
fmt.Println(x) // Still 10 (not modified)
```

**Pass by pointer (reference):**
```go
func updatePointer(ptr *int) {
    *ptr = 100 // Changes the original value
}

x := 10
updatePointer(&x)
fmt.Println(x) // Now 100 (modified!)
```

### Common Use Cases

```go
// 1. Modifying function parameters
func reset(counter *int) {
    *counter = 0
}

// 2. Avoiding large copies (efficiency)
type LargeStruct struct {
    data [1000]int
}

func process(ls *LargeStruct) {
    // Works with original, no copy
}

// 3. nil pointers for optional values
var ptr *int = nil
if ptr == nil {
    fmt.Println("No value set")
}
```

---

## Structs (Like Java Classes)

Structs are Go's way of creating custom types.

### Basic Struct

```go
// Define struct
type Person struct {
    Name string
    Age  int
    City string
}

// Create instance
p1 := Person{
    Name: "Alice",
    Age:  30,
    City: "NYC",
}

// Short form (order matters)
p2 := Person{"Bob", 25, "LA"}

// Access fields
fmt.Println(p1.Name) // Alice
p1.Age = 31          // Modify field
```

### Methods on Structs

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver (doesn't modify original)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can modify original)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area()) // 50

rect.Scale(2)
fmt.Println(rect.Width)  // 20
fmt.Println(rect.Area()) // 200
```

### When to Use Pointer Receivers?

```go
// Use pointer receiver (*Type) when:
// 1. Method needs to modify the struct
// 2. Struct is large (avoid copying)
// 3. For consistency (if one method uses *, use it for all)

// Use value receiver (Type) when:
// 1. Struct is small
// 2. Method doesn't modify data
// 3. Struct is immutable
```

---

## Arrays and Slices

### Arrays (Fixed Size)

```go
// Fixed size, cannot grow
var arr [5]int
arr[0] = 10

// Initialize
numbers := [3]int{1, 2, 3}
```

### Slices (Dynamic, Like Java ArrayList)

```go
// Create slice
var slice []int           // nil slice
slice = []int{1, 2, 3}    // initialized
slice2 := make([]int, 5)  // length 5, all zeros

// Append (IMPORTANT: must assign back!)
slice = append(slice, 4)
slice = append(slice, 5, 6, 7)

// Length and capacity
fmt.Println(len(slice))  // 7
fmt.Println(cap(slice))  // capacity (may be larger)

// Slicing
sub := slice[1:4]  // Elements at index 1, 2, 3
```

### Why Assign Back After Append?

```go
// append() returns a NEW slice
nums := []int{1, 2, 3}
nums = append(nums, 4)  // MUST assign back

// Wrong!
append(nums, 5)  // This does nothing, result is lost
```

---

## Concurrency: Goroutines

Goroutines are lightweight threads managed by Go runtime (not OS threads).

### Basic Goroutine

```go
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Run in new goroutine
    go printNumbers()
    
    // Main continues immediately
    fmt.Println("Main function")
    
    time.Sleep(1 * time.Second) // Wait for goroutine
}
```

### WaitGroup (Proper Synchronization)

```go
import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decreases counter when function returns
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)       // Increase counter BEFORE starting goroutine
        go worker(i, &wg)
    }
    
    wg.Wait() // Block until counter reaches 0
    fmt.Println("All workers completed")
}
```

### Key WaitGroup Rules

1. **`wg.Add(1)`** - Call BEFORE starting goroutine (increases counter)
2. **`defer wg.Done()`** - Put at START of goroutine function (decreases counter on return)
3. **`wg.Wait()`** - In main function, blocks until counter = 0

---

## Channels (Communication Between Goroutines)

Channels are pipes for sending data between goroutines.

### Basic Channel

```go
// Create channel
ch := make(chan int)

// Send data (blocks until received)
go func() {
    ch <- 42  // Send 42 to channel
}()

// Receive data (blocks until sent)
value := <-ch
fmt.Println(value) // 42
```

### Buffered Channels

```go
// Buffered channel (won't block until full)
ch := make(chan int, 2)

ch <- 1  // OK, buffer has space
ch <- 2  // OK, buffer full now
// ch <- 3  // Would block! Buffer is full

fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 2
```

### Closing Channels

```go
ch := make(chan int)

go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch) // Signal no more values
}()

// Receive until closed
for value := range ch {
    fmt.Println(value)
}
```

---

## Mutex (Preventing Race Conditions)

Use mutex when multiple goroutines access shared data.

### Without Mutex (Race Condition)

```go
var counter = 0

func increment() {
    counter++ // NOT safe!
}

// Multiple goroutines cause race condition
for i := 0; i < 1000; i++ {
    go increment()
}
```

### With Mutex (Safe)

```go
import "sync"

var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()         // Acquire lock
    counter++         // Safe now
    mu.Unlock()       // Release lock
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }
    
    wg.Wait()
    fmt.Println(counter) // Always 1000 (correct!)
}
```

### RWMutex (Read-Write Lock)

```go
var (
    data   map[string]string
    rwMu   sync.RWMutex
)

// Multiple readers can access simultaneously
func readData(key string) string {
    rwMu.RLock()         // Read lock
    defer rwMu.RUnlock()
    return data[key]
}

// Only one writer at a time
func writeData(key, value string) {
    rwMu.Lock()          // Write lock (exclusive)
    defer rwMu.Unlock()
    data[key] = value
}
```

**Use RWMutex when:**
- Many reads, few writes
- Read operations are frequent
- Want better performance for concurrent reads

---

## Common Patterns

### Error Handling Pattern

```go
func doSomething() error {
    // Check errors immediately
    result, err := someOperation()
    if err != nil {
        return fmt.Errorf("someOperation failed: %w", err)
    }
    
    // Use result
    return nil
}
```

### Defer for Cleanup

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Executes when function returns
    
    // Work with file
    // file.Close() automatically called on return
    return nil
}
```

### Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2  // Process job
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for a := 1; a <= 9; a++ {
        fmt.Println(<-results)
    }
}
```

---

## Quick Reference

| Java | Go |
|------|-----|
| `class Person { }` | `type Person struct { }` |
| `try-catch` | `if err != nil { }` |
| `public/private` | `Uppercase/lowercase` |
| `void` | No return type |
| `ArrayList<Integer>` | `[]int` (slice) |
| `Thread` | `goroutine` |
| `synchronized` | `sync.Mutex` |
| `null` | `nil` |
| `instanceof` | Type assertion `value.(Type)` |

---

## Tips for Java Developers

1. **No classes** - Use structs with methods
2. **Explicit error handling** - No try-catch, check `err != nil`
3. **Pointers matter** - Use `&` and `*` correctly
4. **Goroutines are cheap** - Create thousands easily
5. **`append()` returns new slice** - Always assign back
6. **Uppercase = public** - Lowercase = package-private
7. **Multiple return values** - Common for `(result, error)`
8. **`defer` is your friend** - Use for cleanup
9. **No inheritance** - Use composition and interfaces
10. **Zero values** - Uninitialized vars have default values (0, "", nil, false)

---

## Next Steps

- Read "Effective Go": https://go.dev/doc/effective_go
- Practice on Go Playground: https://go.dev/play/
- Explore standard library: https://pkg.go.dev/std
- Try building small projects (REST API, CLI tools)
