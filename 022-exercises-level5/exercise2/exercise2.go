package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	newTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	fmt.Println("Truck:", newTruck)
	fmt.Println("Truck doors:", newTruck.doors)

	fmt.Println()

	newSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}
	fmt.Println("Sedan:", newSedan)
	fmt.Println("Sedan doors:", newSedan.doors)
	fmt.Println("Sedan color:", newSedan.color)
	fmt.Println("Sedan luxury:", newSedan.luxury)
}
