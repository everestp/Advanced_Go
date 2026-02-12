package main

import (
	"fmt"
	"time"
)

// func main(){
// 	done := make(chan struct{})

// 	go func(){
// 		fmt.Println("Working")
// 		  time.Sleep(2 * time.Second)
// 		 done <- struct{}{}

// 	}()

// 	<-done
// 	fmt.Println("Finished")
// }

// func main(){
// 	ch := make(chan int)
// 	go func() {
// 	  ch <- 9	 // Blocking until the value is recieved
// 	  fmt.Println("Sent value")
// 	}()

// 	value := <-ch       // Blocking until the value is send
// 	fmt.Println(value)
// }


// func main(){
// 	numGoroutines :=3 
// 	done := make(chan int ,3 )

// 	for  i:= range numGoroutines{
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working ...\n",id)
// 			time.Sleep(time.Second)
// 			done <- id // SENDING SIGNAL OF COMPLETION
// 		}(i)
// 	}
// 	for range numGoroutines{
// 		<-done // Wait for each go routine to finihsed

// 	}
// 	fmt.Println("All goroutine is fineshed")
// }

//======SYNC DATA EXCHANGE=====
func channel_sync(){
   data := make(chan string)
    go func() {
		for i := range 5{
			 data <- "hello" + string(i)
		}
	}()

   for value := range data{
	fmt.Println("Received value", value,":", time.Now() )
   }



}