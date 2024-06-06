package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps ")

	language := make(map[string]string)

	language["C"] = "C++"
	language["G"] = "Go"
	language["J"] = "Java"

	fmt.Println("maps are : ", language)

	fmt.Println("C for => ", language["C"])
	fmt.Println("G for => ", language["G"])
	fmt.Println("J for => ", language["J"])

	//delete(language, "J")

     delete(language,"J")

	fmt.Println(language)

	//loops in maps
        
   for key,value := range language{
	 fmt.Printf("value of key %v value=> %v \n",key,value)
   }

	fmt.Println(language)
}
