package main

import (
	"fmt"
	"time"
)

func non_blocking_operation(){
	// ====NON BLOCKING RECEIVE OPERATION
// 	ch := make(chan int)
// 	select{
// 	case msg := <-ch:
// 	fmt.Println("Recieved",msg)
// 	default:
// 		fmt.Println("No message available")
// 	}

// 	//==NON BLOCKIN SEND OPERATION

// 	select {
// 	case ch <-1:
// 		fmt.Println("Send message")
// 	default:
// 		fmt.Println("Channel is not ready to receive")
// 	}

// == NON BLOCKING OPERATION IN REAL TIME SYSTEM
data := make(chan int)
quite := make(chan bool)
 go func() {
	for {
		select{
			case d := <- data:
			fmt.Println("Data received :",d)
			case <- quite:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data ")
				time.Sleep(500 * time.Millisecond)
		}
	}
 }()

 for i := range 5 {
	data <-i
	time.Sleep(time.Second)
 }
quite<-true

}