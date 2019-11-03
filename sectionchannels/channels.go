package sectionchannels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Receivers always block until there is data to receive.
// If the channel is unbuffered, the sender blocks until the receiver has received the value.
// If the channel has a buffer, the sender blocks only until the value has been copied to the buffer;
// If the buffer is full, this means waiting until some receiver has retrieved a value.

// This usage of channel will result to a deadlock.
//
//func failingChannel() {
//	c := make(chan int)
//	c <- 42
//	fmt.Println(<-c)
//}

func unbufferedChannel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("unbufferedChannel", <-c)
}

func bufferedChannel() {
	c := make(chan int, 1)
	c <- 42
	// c <- 43 // Try to uncomment this line of code.
	fmt.Println("bufferedChannel", <-c)
}

// Directional channels allows you to only read or write with given channel.
// For a send func we are casting our multi dimensional channel as a send only channel.
// For a receive func we are casting our multi dimensional channel as a receive only channel.
// Multi dimensional channels are higher level compared to send/receive channel so it can be
// casted as a send/receive channel but send/receive channels can not be casted a multi dimensional.
func directionalChannels() {
	c := make(chan int)

	// Launching send in a separate Go routine
	// It allows us to run send func as a concurrent func.
	go send(c)

	// Receive func will wait for a send func to write item
	// to the channel.
	receive(c)

}

func send(c chan<- int) {
	value := 42
	c <- value

	fmt.Println("Sent:", value)
}

func receive(c <-chan int) {
	value := <-c
	fmt.Println("Received:", value)
}

func rangeOverChannel() {
	c := make(chan int)

	// Writing values we want and close channel afterwards.
	go sendMultiple(c)

	// Receiving multiple entries from a channel.
	receiveMultiple(c)
}

func sendMultiple(c chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Printf("Wrote %v to the channel. \n", i)
		c <- i
	}

	close(c)
}

func receiveMultiple(c <-chan int) {
	for value := range c {
		fmt.Printf("Received %v from the channel. \n", value)
	}
}

func selectFromChannels() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		quit <- 0
	}()

	for {
		select {
		case v := <-odd:
			fmt.Println("From odd", v)
		case v := <-even:
			fmt.Println("From even", v)
		case <-quit:
			fmt.Println("From quit")
			return
		}
	}
}

func commaOkIdiom() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	// Ok shows if channel is opened. In this case it will return 42 and true.
	v, ok := <-c
	fmt.Println(v, ok)

	// After first write to a channel we closed it so now it will result to 0 and false.
	v, ok = <-c
	fmt.Println(v, ok)
}

// Merge values from several channels into one.
func mergeFewChannels() {
	var messages []string

	c1 := createInfiniteChannel("Nikas")
	c2 := createInfiniteChannel("Naste")
	c := merge(c1, c2)

	// For simplicity reason just waiting for 10 values to emit.
	for i := 0; i < 10; i++ {
		// Read from a channel will block execution of code below.
		// So, you will not see fmt.Println until length of slice
		// will be 10.
		message := <-c
		messages = append(messages, message)
	}

	fmt.Println("All required data gathered.")

	for i, v := range messages {
		fmt.Println(i, v)
	}
}

func createInfiniteChannel(name string) <-chan string {
	c := make(chan string)

	// Start a new Go routine to write message to the channel.
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s message number %d", name, i+1)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

// Fan in
func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup

	out := make(chan string)
	handler := func(c <-chan string) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	// Ensures that all Go routines will be finished.
	wg.Add(len(cs))

	// Starting separate Go routines for each channel.
	// It allows us to merge multiple channels.
	for _, c := range cs {
		go handler(c)
	}

	// Starting another Go routine to wait for a data
	// from a channels and closing up out channel afterwards.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func parallelTimeConsumingWork() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Populating our 1st channel with data for heavy operations.
	go populate(c1)

	// Processing our populated channel with long running tasks.
	go fanOutIn(c1, c2)

	// Response after one of long running tasks finish.
	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("About to exit.")
}

func populate(c chan<- string) {
	for i := 0; i < 100; i++ {
		c <- fmt.Sprintf("Video to convert nr %v", i)
	}
	close(c)
}

func fanOutIn(sc <-chan string, rc chan<- string) {
	var wg sync.WaitGroup

	for v := range sc {
		wg.Add(1)
		go func(v2 string) {
			rc <- processLongRunningTask(v)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(rc)
}

func fanOutInWithThrottle(sc <-chan string, rc chan<- string) {
	const delta = 10
	var wg sync.WaitGroup

	wg.Add(delta)

	for i := 0; i < delta; i++ {
		go func() {
			for v := range sc {
				func(v2 string) {
					rc <- processLongRunningTask(v)
				}(v)
				wg.Done()
			}
		}()
	}

	wg.Wait()
	close(rc)
}

func processLongRunningTask(task string) string {
	time.Sleep(time.Duration(rand.Intn(5e3)) * time.Millisecond)
	return fmt.Sprintf("Finished task - %v", task)
}

func Example() {
	//unbufferedChannel()
	//bufferedChannel()
	//directionalChannels()
	//rangeOverChannel()
	//selectFromChannels()
	//commaOkIdiom()
	//mergeFewChannels()
	parallelTimeConsumingWork()
}
