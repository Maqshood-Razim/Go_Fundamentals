package main

import "fmt"

func main() {
	fmt.Println("Structs...")

	raz := User{"Razim", "Razim@gmail.com", true, 20}
	fmt.Println(raz)
	fmt.Printf("user is %+v\n",raz)
	fmt.Printf("name is %v age is %v",raz.Name,raz.Age)
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
