package main

import "fmt"

func main() {
	fmt.Println("closures")

	newcounter := makecounter()
	fmt.Println(newcounter())
	fmt.Println(newcounter())
	fmt.Println(newcounter())

}

func makecounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
