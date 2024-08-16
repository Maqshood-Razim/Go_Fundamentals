package main

import "fmt"

func main() {
	fmt.Println("Runes")


	var r rune =  'a'

	fmt.Println(r)


	s := "hello,world"


	for _,r := range s{
		fmt.Printf("runes are %c\n ",r)
     	}

   // converts string to rune
	var h string = "hello,world"

	runes := []rune(h)

	fmt.Println(runes)

	//converts rune to string


	runesslice := []rune{'h','e','l','l','o'}

    t := string(runesslice)

	fmt.Println(t)
	
		
}
