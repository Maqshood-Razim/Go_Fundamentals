package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		// time.Sleep(700 * time.Millisecond)
	}
}

func printhello(){
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
		// time.Sleep(700 *time.Millisecond)
	}
}

func main() {
	// Start printNumbers in a new goroutine
	go printNumbers()

	// Run printNumbers in the main goroutine
	printNumbers()
     
	go printhello()

	go printNumbers()



	// Wait for goroutines to finish (not necessary if main goroutine blocks)
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
