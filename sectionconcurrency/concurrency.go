package sectionconcurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const delta = 100

func mutexExample() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter = 0

	fmt.Println("Number of cpu:", runtime.NumCPU())
	fmt.Println("Number of Go routines:", runtime.NumGoroutine())

	wg.Add(delta)

	for i := 0; i < delta; i++ {
		go func() {
			mu.Lock()

			value := counter
			runtime.Gosched()
			value++
			counter = value
			fmt.Println("Counter:", counter)

			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

func atomicExample() {
	var wg sync.WaitGroup
	var counter int64 = 0

	fmt.Println("Number of cpu:", runtime.NumCPU())
	fmt.Println("Number of Go routines:", runtime.NumGoroutine())

	wg.Add(delta)

	for i := 0; i < delta; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

// Example is function which shows how to use concurrency in Go.
func Example() {
	start := time.Now()

	mutexExample()
	t := time.Now()
	fmt.Println(t.Sub(start).Seconds())

	start = time.Now()
	atomicExample()
	t = time.Now()
	fmt.Println(t.Sub(start).Seconds())
}
