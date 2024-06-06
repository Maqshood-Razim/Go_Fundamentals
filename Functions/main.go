package main

import "fmt"

func main() {
	fmt.Println("Functions")

	greeter()
	greetertwo()

	// result := adder(19, 7)
	// fmt.Println("result is : ", result)

	fmt.Println(adder(19,7))

	prores := proadder(19, 3, 7, 77)

	fmt.Println("result is : ", prores)

	multres,messege := multiple(7, 3, 19)

	fmt.Println("result of multiplication : ", multres)
	fmt.Println("messege is : ",messege)

}

func greetertwo() {
	fmt.Println("i'm greeter two")
}

func greeter() {
	fmt.Println("welcome")
}

func adder(a int, b int) int {
	return a * b
}

func proadder(values ...int) int {
	total := 0

	for _, val := range values {
		total = total + val
	}

	return total
}

func multiple(valuesmultiple ...int) (int,string) {

	resultofmult := 1

	for _, res := range valuesmultiple {
		resultofmult = resultofmult * res
	}

	return resultofmult,"this is a multiplication"
}
