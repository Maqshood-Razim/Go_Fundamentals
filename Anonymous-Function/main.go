package main

import (
	"fmt"

	
)

func main() {

	fmt.Println("Anonymous Function")

	func (){
		fmt.Println("this is a anonymous function")
	}()


	//anonymous function with parameters


	func (val string)  {
	  fmt.Println(val)
	}("this is a messege")


	//anonymous function with return values

    result := func (a int, b int) int  {
		  return a+b   
	}(19,7)

	fmt.Println("result is : ",result)



	// assingning anonymous function in a variable


	multiple := func (valone , valtwo int) int {
		   return valone*valtwo
	} 

	
	fmt.Println(multiple(7,3))

}
