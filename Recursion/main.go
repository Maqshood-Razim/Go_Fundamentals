package main

import "fmt"

func main() {
	fmt.Println("recursion")

	fmt.Println(factorial(7))

}

func factorial(n int) int {

	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)

}
