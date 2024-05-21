package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating ", input)
	
	numrating,err := strconv.ParseFloat(strings.TrimSpace(input),64)


	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("added number ",numrating+1)
	}

//      var d int

//    fmt.Scan(&d)

//    fmt.Println("entered number is ",d)

//    fmt.Println("added number is",d+1)



 
}
