package main

import (
	"fmt"
	"time"
)

// Future Reference: The Select Statement
// 'select' blocks until one of its cases can run, then it executes that case. 
// If multiple cases are ready, it chooses one pseudo-randomly. This prevents 
// starvation where one channel always dominates the others.

func multiplexoing_select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Logic Block: Simultaneous Data Production
	// We simulate two different services returning data at different speeds.
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Service A Response"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Service B Response"
	}()

	// Logic Block: The Multi-Channel Listener
	// We use a loop with select to keep listening until we get what we need.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Result:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Result:", msg2)
		case <-time.After(3 * time.Second):
			// Future Reference: Timeouts with time.After
			// This case creates a timer channel. If no data arrives within the
			// duration, this case triggers, preventing the program from hanging.
			fmt.Println("Error: Operation timed out!")
			return
		}
	}

	fmt.Println("--- Switching to Closure Check ---")
	closureExample()
}

// Future Reference: The Comma-OK Idiom in Select
// When receiving 'msg, ok := <-ch', 'ok' is false if the channel is empty 
// AND closed. It is critical to check 'ok' to avoid infinite loops 
// reading zero-values from a closed channel.

func closureExample() {
	jobs := make(chan int)

	go func() {
		jobs <- 101
		close(jobs) // Signalling no more work exists
	}()

	for {
		select {
		case msg, ok := <-jobs:
			if !ok {
				// Logic Block: Cleanup and Exit
				// Once 'ok' is false, we break or return to stop the loop.
				fmt.Println("Channel closed. Cleaning up resources...")
				return
			}
			fmt.Printf("Processing Job: %d\n", msg)
		}
	}
}
