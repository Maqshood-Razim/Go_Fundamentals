package main

import "fmt"

const Logged = 191919 // this is public

func main() {
	// var str string = "Ronaldo"
	// fmt.Println(str)
	// var (
	// 	messi  string = " messi"
	// 	neymar string = "neymar"
	// )
	// fmt.Println(messi, neymar)

	var username string = "razim"
	fmt.Println(username)
	fmt.Printf("varible type is %T \n", username)
	var username1 string = "razim t"
	fmt.Println(username1)

	//boolean

	var islogged bool = true
	fmt.Printf("variable type is %T \n", islogged)
	fmt.Println(islogged, Logged)

	// numbers

	// var smallval uint = 219
	// fmt.Printf("variable type is %T ", smallval)
	// fmt.Println(smallval)

	// var smallvalfloat float64 = 219.19191919
	// fmt.Printf("variable type is %T ", smallvalfloat)
	// fmt.Println(smallvalfloat)

	// implicit type

	var name = "maqshood"
	fmt.Println(name)

	// no var type or operater

	numberofusers := 19999.9
	fmt.Println("varles operater => ", +numberofusers)

	fmt.Println(Logged)

	name2 := "razim"
	fmt.Println(name, name2)

	// if else

	number := 19

	var result string

	if number > 19 {
		result = "number is big"
	} else if number < 19 {
		result = "number is small"
	} else {
		result = "number is correct"
	}

	fmt.Println(result)

	if 19%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	

}
