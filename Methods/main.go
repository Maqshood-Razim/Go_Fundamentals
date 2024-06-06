package main

import "fmt"

func main() {
	fmt.Println("Methods")

	raz := User{"Razim", "Razim@gmail.com", 20}
	fmt.Println(raz)
	fmt.Printf("user is %+v\n", raz)
	fmt.Printf("name is %v age is %v\n", raz.Name, raz.Age)
	raz.NewMail()
	fmt.Println(raz)

}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("new mail is : ", u.Email)
}
