package main

import "fmt"

// "math"

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
// 	fmt.Printf("total milege per month is %.2f km/l" , tm)
// }

// another example

// type Shape interface {
// 	Area() float32
// 	perimeter() float32
// }

// type Circle struct {
// 	Radious float32
// }

// func (c Circle) Area() float32 {
// 	return math.Pi * c.Radious * c.Radious
// }

// func (c Circle) perimeter() float32 {
// 	return 2 * math.Pi * c.Radious
// }

// func printshape(s Shape) {
// 	fmt.Printf("area : %.2f\n",s.Area())
// 	fmt.Printf("perimeter : %.2f",s.perimeter())
// }

type speaker interface {
	speak() string
}

type mover interface {
	move() string
}

type robot struct {
	model string
}

type cat struct {
	name string
}

type car struct {
	name string
}

func (r robot) speak() string {
	return "beep"
}

func (r robot) move() string {
	return "rolling wheels"
}

func (c cat) speak() string {
	return "mew"
}
func (c cat) move() string {
	return "walkning"
}

func (ca car) speak() string {
	return "...."
}

func (ca car) move() string {
	return "on wheels"
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

	// res := Circle{
	// 	Radious: 7,
	// }

	// printshape(res)

	// type Assertion

	var i interface{}="oi"
	 p:= i.(string)
	 fmt.Println(p)

	//  //type Switch

	// switch v:=i.(type){
	// case string:
	// 	fmt.Println("string:",v)
	// case int:
	// 	fmt.Println("int:",v)
	// default:
	// 	fmt.Println("unkown")
	// }

	var s speaker
	var m mover

	r := robot{model: "rd3"}
	c := cat{name: "anything"}
	ca := car{name: "TATA"}

	s = r
	fmt.Println(s.speak())

	m = r
	fmt.Println(m.move())

	s = c
	fmt.Println(s.speak())

	m = c

	fmt.Println(m.move())

	cm := ca
	fmt.Println(cm.speak())

	cs := ca

	fmt.Println(cs.move())

}
