package main

import (
	"fmt"
	
	"time"
)

//GOroutine are just the function that leave the main thread and run in the background and come back join th  main thread once the function are finished/ ready to return value
// Goroutine do not stop the program flow and are non blocking


func main(){
	var err error
	fmt.Println("Program start")
	go func ()  {
		err =  doWork()
	}()
  go sayHello()
 // err =  doWork() this is  not accepted
  go printLetters()
  go printNumbers()
   if err != nil {
	fmt.Println("Errpor", err)
   } else{
	 fmt.Println("work completed sucessfully")
   }
  time.Sleep(2 * time.Second)
}


func sayHello(){
	time.Sleep(1 *time.Second)
	fmt.Println("Hello from goroutine")
	
}

func  printNumbers(){
	 for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)

	 }
}



func printLetters(){
	for _ ,letter := range "abcd"{
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error{
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Error occured to do work")
}