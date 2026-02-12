package main

import (
	"fmt"
	"time"
)


func main(){
	ch := make(chan int)
	producer(ch)
	consumer(ch)
    // go func(ch chan <-int) {
	// 	for i := range 5 {
	// 		ch <-i
	// 	}
	// 	close(ch)
	// }(ch)

	// for  value := range ch {
	// 	fmt.Println("Received",value)
	// }
	revceivedData(ch)
	
	time.Sleep(time.Second)
}

//Recieved only cahnnel 

func revceivedData(ch <- chan int){
	for  value := range ch {
		fmt.Println("Received only",value)
	}
}
func consumer(ch <- chan int){
	for  value := range ch {
		fmt.Println("Received only Consumer",value)
	}
}
func producer(ch chan <- int){
	 go func(ch chan <-int) {
		for i := range 5 {
			ch <-i
		}
		close(ch)
	}(ch)
}