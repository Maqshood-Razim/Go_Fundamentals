package main

import "fmt"

func main() {

	fmt.Println("Array")

	var fruitarray [5] string

	fruitarray[0]="apple"
	fruitarray[1]="dates"
	fruitarray[3]="mango"
	fruitarray[4]="watermelon"

	fmt.Println("array is \n",fruitarray)
	fmt.Println("length is : ",len(fruitarray))


	//another method of declaring array

	 veg := [3] string {"beans","mushroom","betroot"}

	fmt.Println(veg)
}