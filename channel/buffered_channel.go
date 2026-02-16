package main

import (
	"fmt"
	"time"
)

// Future Reference: Buffered vs. Unbuffered Blocking
// Unbuffered: Sender blocks until Receiver is ready (Handshake).
// Buffered: Sender only blocks when the buffer is FULL.
// Receiver only blocks when the buffer is EMPTY.

func main() {
	// Logic Block: Initialization
	// We create a channel with a capacity of 2.
	// This means we can send 2 items without needing a concurrent receiver.
	ch := make(chan int, 2)

	// Logic Block: Non-blocking Sends
	// These two lines execute immediately because the buffer has room.
	ch <- 10
	ch <- 20
	fmt.Println("Buffer is now FULL (2/2). No blocking occurred yet.")

	// Logic Block: Asynchronous Drain
	// We spin up a goroutine to remove one item after a delay.
	// This will eventually "free up" a slot in the buffer.
	go func() {
		fmt.Println("Worker: Sleeping for 2 seconds...")
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker: Received %d. Buffer now has 1 slot free.\n", <-ch)
	}()

	// Logic Block: The Blocking Send
	// This line will HALT execution. Since the buffer is 2/2, 
	// this send must wait for the goroutine above to read from the channel.
	fmt.Println("Main: Attempting to send 3rd value (this will block)...")
	ch <- 30 
	
	fmt.Println("Main: Successfully sent 30 after worker freed a slot!")
	
	// Logic Block: Final Drain
	fmt.Println("Received remaining:", <-ch)
	fmt.Println("Received remaining:", <-ch)
}