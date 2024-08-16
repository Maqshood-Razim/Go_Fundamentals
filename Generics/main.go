package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("Generics")
	fmt.Println(add(3.19,7))


}
func add[T constraints.Ordered] (a T,b T) T{
	return a+b
}