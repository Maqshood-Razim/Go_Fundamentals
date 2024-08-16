package main

import "fmt"

func main() {
	fmt.Println("Structs...")

	// raz := User{"Razim", "Razim@gmail.com", 20}
	// fmt.Println(raz)
	// fmt.Printf("user is %+v\n", raz)
	// fmt.Printf("name is %v age is %v", raz.Name, raz.Age)

	// acc := Account{"Razim", "Razim@gmai.com", 19}

	// fmt.Printf("account details %+v \n", acc)

	// fmt.Printf("%v %v", acc.Name, acc.Age)

	p := Person{Name: "razim", Age: 20, Email: "razim@gmail.com"}

	p1 := Person{Name: "maq", Age: 23, Email: "maq@gmail.com"}

	fmt.Println(p.greet())

	p.birthday()

	fmt.Println("age : ", p.Age)

	greeting := p1.greet()
	fmt.Println(greeting)

    p1.havebirthday()

}

func (p Person) greet() string {
	return "the name is " + p.Name
}

func (p *Person) birthday() {
	p.Age++
}


func (p *Person) havebirthday(){
	p.Age++
}

// type User struct {
// 	Name   string
// 	Email  string
// 	Age    int
// }

// type Account struct {
// 	Name  string
// 	Email string
// 	Age   int
// }

type Person struct {
	Name  string
	Age   int
	Email string
}
