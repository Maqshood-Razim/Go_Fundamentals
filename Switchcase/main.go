package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	rand.New(rand.NewSource(time.Now().UnixNano()))
	dicenumber := rand.Intn(7)

	fmt.Println(dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("you got 1 you can move ")
	case 2:
		fmt.Println("you got 2 you can move 2 spots ")
	case 3:
		fmt.Println("you got 3 you can move 3 spots ")
	case 4:
		fmt.Println("you got 4 you can move 4 spots ")
	case 5:
		fmt.Println("you got 5 you can move 5 spots ")
	case 6:
		fmt.Println("you got 6 you can move 6 spots ")
	default:
		fmt.Println("what was that!")
	}
}
