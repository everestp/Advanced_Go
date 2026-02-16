package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Future Reference: The "Empty Struct" Signal
// Using 'chan struct{}' is a memory-efficient way to send signals. 
// Since struct{} occupies 0 bytes, it's the idiomatic choice when 
// you don't care about the data, only that "something happened."

func signalPattern() {
	done := make(chan struct{})

	go func() {
		// Logic Block: Simulate a background task.
		fmt.Println("Worker: Starting task...")
		time.Sleep(1 * time.Second)
		
		// Signal completion by sending an empty struct.
		done <- struct{}{}
	}()

	// Logic Block: Block execution until the signal is received.
	<-done
	fmt.Println("Main: Signal received, moving on.")
}

// Future Reference: sync.WaitGroup vs. Channels for Sync
// While channels can synchronize goroutines, sync.WaitGroup is 
// specifically designed for "Wait for All" scenarios. It is more 
// readable and performant when tracking multiple concurrent tasks.

func multiWorkerPattern() {
	var wg sync.WaitGroup
	numWorkers := 3

	for i := range numWorkers {
		// Increment the counter for each goroutine launched.
		wg.Add(1)
		
		go func(id int) {
			// Logic Block: Ensure Done() is called even if the function panics.
			defer wg.Done()
			
			fmt.Printf("Worker %d: Processing...\n", id)
			time.Sleep(500 * time.Millisecond)
		}(i)
	}

	// Logic Block: Block until the WaitGroup counter hits zero.
	wg.Wait()
	fmt.Println("All workers finished successfully.")
}

// Future Reference: Avoiding Channel Deadlocks
// A 'for range' loop on a channel will continue until the channel 
// is closed. If the sender never calls 'close()', the receiver 
// hangs forever, triggering a deadlock panic.



func dataExchangePattern() {
	dataCh := make(chan string)

	go func() {
		// Logic Block: Clean up by closing the channel when production finishes.
		defer close(dataCh)

		for i := range 3 {
			// Using strconv.Itoa is necessary to convert int to "string" digits.
			dataCh <- "Payload " + strconv.Itoa(i)
		}
	}()

	// Logic Block: Iteratively drain the channel.
	for val := range dataCh {
		fmt.Println("Received:", val)
	}
}

func channel_sync() {
	fmt.Println("--- Signal Pattern ---")
	signalPattern()

	fmt.Println("\n--- Multi-Worker Pattern ---")
	multiWorkerPattern()

	fmt.Println("\n--- Data Exchange Pattern ---")
	dataExchangePattern()
}