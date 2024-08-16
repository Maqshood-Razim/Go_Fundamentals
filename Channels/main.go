package main

import "fmt"

func main() {
	fmt.Println("Channles")

	ch := make(chan int)

	go func() {
		ch <- 19
	}()

	value := <-ch
	fmt.Println(value)

	//buffered channles

	chb := make(chan int, 3)

	chb <- 1
	chb <- 2
	chb <- 3
	fmt.Println(<-chb)
	fmt.Println(<-chb)
	fmt.Println(<-chb)


	done := make(chan bool, 1)

	go sayHello(done)

	<-done // Wait for the goroutine to finish
	fmt.Println("Goroutine finished")

}
func sayHello(done chan bool) {
	fmt.Println("Hello from goroutine!")
	done <- true // Signal that the goroutine is done
}
