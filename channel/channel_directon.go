package main

import (
	"fmt"
	"sync"
)

// Future Reference: Channel Directionality & Safety
// By specifying 'chan<-' (send-only) or '<-chan' (receive-only), we enforce 
// compile-time checks. This prevents a consumer from accidentally closing 
// a channel or a producer from trying to read from one.

func channel_direction() {
	// Initialize an unbuffered channel for synchronous communication.
	ch := make(chan int)
    
	// Use a WaitGroup to ensure the main function waits for the consumer to finish.
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the producer. It runs its own goroutine internally.
	producer(ch)

	// Start the consumer in a separate goroutine to allow concurrent processing.
	go consumer(ch, &wg)

	// Block main until wg.Done() is called inside the consumer.
	wg.Wait()
}

// Future Reference: The "Close-by-Producer" Pattern
// In Go, the sender should always be the one to close the channel. 
// This signals to the 'range' loop in the consumer that no more data 
// is coming, preventing the "all goroutines are asleep - deadlock" error.

func producer(ch chan<- int) {
	// Launch a background worker to prevent the main thread from blocking on send.
	go func() {
		for i := range 5 {
			fmt.Printf("Producing: %d\n", i)
			ch <- i
		}
		// Closing the channel is vital for the consumer's 'for-range' loop to exit.
		close(ch)
	}()
}

// Future Reference: Synchronizing with WaitGroups
// When a consumer runs in a goroutine, 'main' might finish before the consumer 
// processes the data. Passing a pointer to sync.WaitGroup allows the consumer 
// to signal back to 'main' that its work is complete.

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	// Ensure the WaitGroup counter is decremented when the function returns.
	defer wg.Done()

	// The 'range' keyword elegantly handles channel draining and exits on close.
	for value := range ch {
		fmt.Printf("Consumed: %d\n", value)
	}
}