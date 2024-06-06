package main

import "fmt"

func main() {
	fmt.Println("Structs...")

	// raz := User{"Razim", "Razim@gmail.com", 20}
	// fmt.Println(raz)
	// fmt.Printf("user is %+v\n", raz)
	// fmt.Printf("name is %v age is %v", raz.Name, raz.Age)

	acc := Account{"Razim", "Razim@gmai.com", 19}

	fmt.Printf("account details %+v \n", acc)

	fmt.Printf("%v %v", acc.Name, acc.Age)

}

// type User struct {
// 	Name   string
// 	Email  string
// 	Age    int
// }

type Account struct {
	Name  string
	Email string
	Age   int
}
