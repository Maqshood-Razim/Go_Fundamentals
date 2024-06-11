package main

import (
	"fmt"
	"math"
)

// type Vehicle interface {
// 	Milege() float32
// }

// type Bmw struct {
// 	Distance     float32
// 	Fuel         float32
// 	AveregeSpeed float32
// }

// type Audi struct {
// 	Distance     float32
// 	Fuel         float32
// 	AveregeSpeed float32
// }

// func (b Bmw) Milege() float32 {
// 	return b.Distance / b.Fuel
// }

// func (a Audi) Milege() float32 {
// 	return a.Distance / a.Fuel
// }

// func totalMilege(m []Vehicle) {
// 	tm := 0.0
// 	for _, val := range m {
//         tm = tm + float64(val.Milege())
// 	}
// 	fmt.Printf("total milege per month is %f km/l" , tm)
// }

// another example

type Shape interface {
	Area() float32
	perimeter() float32
}

type Circle struct {
	Radious float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radious * c.Radious
}

func (c Circle) perimeter() float32 {
	return 2 * math.Pi * c.Radious
}

func printshape(s Shape) {
	fmt.Printf("area : %.2f\n",s.Area())
	fmt.Printf("perimeter : %.2f",s.perimeter())
}


func main() {

	// b1 := Bmw{
	// 	  Distance: 199.2,
	// 	  Fuel: 19.7,
	// 	  AveregeSpeed: 99.7,
	// }

	// a1 := Audi{
	//  Distance: 219.3,
	//  Fuel: 19.3,
	//  AveregeSpeed: 97.3,
	// }

	// person := []Vehicle{b1,a1}
	// totalMilege(person)

	res := Circle{
		Radious: 7,
	}

	printshape(res)

}
