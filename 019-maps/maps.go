package main

import "fmt"

func main() {

	coords := map[string]int{
		"x": 10229222,
		"y": 82732711,
	}
	fmt.Println(coords)

	// Getting element's map
	fmt.Println(coords["x"])

	// Getting element's unexisting in map
	fmt.Println(coords["z"])

	// Avoid unexisting elements
	if v, ok := coords["z"]; ok {
		fmt.Println("ok value: ", ok) // => ok value is false
		fmt.Println("Printing existing values: ", v)
	}

	if v, ok := coords["x"]; ok {
		fmt.Println("ok value: ", ok)
		fmt.Println("Printing existing values: ", v)
	}

	// Empty Map
	newCoords := map[string]int{}
	fmt.Println(newCoords)
	//Adding elements
	newCoords["z"] = 123781282
	fmt.Println(newCoords)
	fmt.Println(newCoords["z"])
	fmt.Println(newCoords["Z"])
}
