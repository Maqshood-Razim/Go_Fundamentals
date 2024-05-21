package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitlist = []string{"apple", "mango", "dates"}

	fmt.Println(fruitlist)

	var fruitlistnew = append(fruitlist, "watermelon", "grapes")

	fmt.Println("new append slice is : ", fruitlistnew)

	var cutfruitlistnew = append(fruitlistnew[:2])

	fmt.Println("new slice = ", cutfruitlistnew)

	highscores := make([]int, 3)

	highscores[0] = 19
	highscores[1] = 23
	highscores[2] = 27
	// highscores [4] = 19

	highscores = append(highscores, 40, 20, 99, 70)

	fmt.Println("highscores are : ", highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

	sort.Ints(highscores)

	fmt.Println("sorted slice =", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// remove value in slice based on index

	courses := []string{"Go", "C", "C++", "Java", "Assembly"}

	var index int = 3

	var coursesdel = append(courses[:index], courses[index+1:]...)

	fmt.Println("result = ", coursesdel)

	

}
