package main

import (
	"fmt"
	"time"
)

// Defining struct
type car struct {
	name  string
	color string
	price int
	//given below a different way to defining field of complex type
	//modelOfCar carModel
	carModel
}

type carModel struct {
	year  time.Time
	model string
}

func main() {

	var kia car
	kia.name = "seltos"
	kia.price = 1200000
	kia.color = "black"
	// kia.modelOfCar = carModel{year: time.Now().UTC(), model: "K1"}
	kia.carModel = carModel{year: time.Now().UTC(), model: "K1"}

	kia.print()

	//initialize and declare struct
	s1, _ := time.Parse(time.RFC3339, time.Now().String())
	toyota := car{
		name:  "Innova",
		color: "White",
		price: 2500000,
		// another way of assigning
		// modelOfCar: carModel{
		// 	year:  s1,
		// 	model: "G",
		// },
		carModel: carModel{
			year:  s1,
			model: "G",
		},
	}
	//update struct value
	toyota.color = "Blue"
	toyota.updateName("glanza")
	toyota.print()
}

func (c car) print() {
	fmt.Printf("%+v\n", c)
}

// NOTE: the go will automatically infer if we pass even the value type in the below receiver function's argument
// to change the pointer to a value use (*)
// e.g  pointerToCar := $toyota --> *pointerToCar will yeild the value i.e toyota
// to change from value to a pointer use (&)
// e.g pointerToCar := $toyota
func (c *car) updateName(name string) {
	c.name = name
}
