package main

import (
	"fmt"
	"time"
)



func main() {

	fmt.Println("handling time")

	presenttime:=time.Now()

	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 Mondayś"))


}
