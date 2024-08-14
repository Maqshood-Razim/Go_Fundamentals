package main

import "fmt"

func main() {
	fmt.Println("variadic funtion")

	
    
	fmt.Println(sum(7,3,20))


   

}

func sum(num ...int) int {
    fmt.Println(num,"")

	total := 0
   
 for _,result := range num{
	total = total+result
 }

 return total
	
}
