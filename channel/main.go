package main

import (
	"fmt"
	"time"
)

func main(){
	 greeting := make(chan string)
	 greetString := "Hello"

	 go func() {
		 greeting <- greetString 
		 greeting <-"Hey Everest"
		 for  _ , e := range "abcde"{
			greeting <- "Aphabet" +" " + string(e)
		 }
	 }()
	 //  greeting <- greetString // blocking because it is contineously tryinh to recieve value.m it is ready to  receive the contineous flow of   data

	 receiver := <- greeting
	 fmt.Println(receiver)
	 receiver = <- greeting

	 fmt.Println(receiver)
	 for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	 }
	 time.Sleep(2 * time.Second)
}