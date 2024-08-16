package main

import "fmt"

func main() {
	fmt.Println("embedding")

    e := job{
		person: person{
           Name: "razim",
		   Age: 20,
		},
		contact:contact{
          Number: 9999999,
		},
		position: "developer",
		salary: 700000,
	}

	fmt.Println("name",e.Name)
	fmt.Println("age",e.Age)
	fmt.Println("position",e.position)
	fmt.Println("salary",e.salary)
	fmt.Println("number",e.Number)

	e.pes()

}

type person struct {
	Name string
	Age  int
}

type job struct {
	person
	contact

	position string
	salary   int
}


type contact struct{
	Number int
}

func (p *person) pes() {
	fmt.Printf("my name is %s age is %d ",p.Name,p.Age)
}
