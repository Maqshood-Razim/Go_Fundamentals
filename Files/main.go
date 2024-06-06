package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "this needs to go to a file"

	file, err := os.Create("./myGofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is : ",length)
	defer file.Close()



   var a int =7
   var b int =9
   var d int
  fmt.Println("a",a)
  fmt.Println("b",b)
  d  =a
  a =b
  b = d

   
   fmt.Println("a",a)
   fmt.Println("b",b)

}
