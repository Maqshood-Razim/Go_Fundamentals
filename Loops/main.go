package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"monday", "wendensday", "thursday", "friday", "saturdasy"}

	fmt.Println(days)

	//   for i :=0; i< len(days);i++{
	// 	fmt.Println(days[i])
	//   }

	// for i:= range days{
	//     fmt.Println(days[i])
	// }

	for index, days := range days {
		fmt.Printf("index is %v value is %v \n", index, days)
	}


	value :=1

   for value<10{
     
	// if value==7 {
	// 	break
	// }


	if value==3 {
		value++
		continue
	}

	if value==2 {
		goto i
	}
	 
	fmt.Println("value is = ",value)
	value++
   }

   
	// if value==2 {
	// 	goto i
	// }
	 
   
   i:
   println("goto worked")
}
