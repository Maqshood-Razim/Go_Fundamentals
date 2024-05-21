package main

import (
	"fmt"
)

func main() {

	welcome := "Welcome to userinput"

	fmt.Println(welcome)

	
	var d string
	//reader:=bufio.NewReader(os.Stdin)

	println("enter a number")
	fmt.Scan(&d)
	// input,_:=reader.ReadString('\n')

	println("entered number is =", d)
	fmt.Printf("type is %T \n", d)

	if d == "19" {
		fmt.Println("number")
	} else {
		fmt.Println(d)
	}

}
