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

	chb := make(chan int, 2)

	chb <- 1
	chb <- 2
	fmt.Println(<-chb)
	fmt.Println(<-chb)

}
