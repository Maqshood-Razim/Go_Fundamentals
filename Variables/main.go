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

	// var username string = "razim"
	// fmt.Println(username)
	// fmt.Printf("varible type is %T \n", username)
	// var username1 string = "razim t"
	// fmt.Println(username1)

	// //boolean

	// var islogged bool = true
	// fmt.Printf("variable type is %T \n", islogged)
	// fmt.Println(islogged, Logged)

	// // numbers

	// var smallval uint = 219
	// fmt.Printf("variable type is %T ", smallval)
	// fmt.Println(smallval)

	// var smallvalfloat float64 = 219.19191919
	// fmt.Printf("variable type is %T ", smallvalfloat)
	// fmt.Println(smallvalfloat)

	// // implicit type

	// var name = "maqshood"
	// fmt.Println(name)

	// // no var type or varless operater

	// numberofusers := 19999.9
	// fmt.Println("varles operater => ", numberofusers)

	// fmt.Println(Logged)

	// name2 := "razim"
	// fmt.Println(name, name2)

	// // if else

	// number := 19

	// var result string

	// if number > 19 {
	// 	result = "number is big"
	// } else if number < 19 {
	// 	result = "number is small"
	// } else {
	// 	result = "number is correct"
	// }

	// fmt.Println(result)

	// if 19%2 == 0 {

	// 	fmt.Println("even")
	// } else {
	// 	fmt.Println("odd")
	// }

	var str string = "ronaldo"
	fmt.Println(str)

	var (
		fas string = "fasal"
		sah string = "sahad"
	)

	fmt.Println(fas, sah)

	var username string = "razim"

	fmt.Printf("variable type is %T\n", username)

	var bo bool = true
	fmt.Println(bo)
	fmt.Printf("variable type is => %T\n", bo)

	var in uint16 = 19191
	var bint uint32 = 1919191919

	fmt.Println(in)
	fmt.Println(bint)

	var sflo float32 = 19.19191919191919191919
	var bflo float64 = 19.1919191919191919191919191919191919

	fmt.Println(sflo)
	fmt.Println(bflo)

	var name = "maqshood"

	var age = 20

	fmt.Println(name, age)

	user := "razim"
	fmt.Println(user)

	n := 19

	if n < 19 {
		fmt.Println("number is less")
	} else if n > 19 {
		fmt.Println("number is big")
	} else {
		fmt.Println("number is correct")
	}

	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

}
