package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Error handling in Go")

	var number int

	fmt.Println("enter a number")

	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("only nummber could be enter")
		fmt.Println(err)
	} else {
		fmt.Println("entered number is  : ", number)

	}

	val, err2 := add(77)

	if err2 != nil {
		fmt.Println("error detected")
		fmt.Println(err2)
	} else {
		fmt.Println("value is : ", val)
	}

}

func add(a int) (int, error) {
	if a == 19 {
		return -1, errors.New("number is ninteen")
	} else {
		return a + 19, nil
	}
}
