package main

import (
	"fmt"
	"time"

	
)

//GOroutine are just the function that leave the main thread and run in the background and come back join th  main thread once the function are finished/ ready to return value
// Goroutine do not stop the program flow and are non blocking


func main(){
  go sayHello()
  time.Sleep(2 * time.Second)
}


func sayHello(){
	time.Sleep(1 *time.Second)
	fmt.Println("Hello from goroutine")
	
}