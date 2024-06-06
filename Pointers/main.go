package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int

	// fmt.Println("value of pointer",ptr)

	mynumber :=19

	var ptr = &mynumber


	fmt.Println("actual value of pointer",ptr)
    fmt.Println("actual value of pointer",*ptr)

	*ptr = *ptr +2

	fmt.Println("new value",mynumber)
}