package main

import (
	"fmt"
	"time"
)

// func main(){
// ch := make(chan int ,2)	
//   ch <- 1
//   go func() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Recieved" ,<-ch)
//   }()
//   fmt.Println("Buffered channel")
 
// }

func main(){
	// ====================BLOCKING ON RECIEVER ONLY IF BUFFER IS  EMPTY
	ch := make(chan int ,2 )
   ch <-1
   ch <-2
	go func ()  {
		time.Sleep(2 * time.Second)
	
		fmt.Println("Recieved", <-ch)
	}()
	ch <-3
	fmt.Println("Value", <-ch)
	fmt.Println("End of program")
}