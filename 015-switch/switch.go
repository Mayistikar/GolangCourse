package main

import "fmt"

func main() {

	/* VALUE CONDITION */
	fmt.Println("VALUE CONDITION")
	value := "Condition"
	switch value {
	case "condition":
		fmt.Println("This not works!!")
	case "Condition":
		fmt.Println("THis works!")
	default:
		fmt.Println("This not works")
	}

	fmt.Println()

	/* FALLTHROUGH SKIPPING */
	fmt.Println("FALLTHROUGH SKIPPING")
	switch {
	case true:
		fmt.Println("Works Here because default value condition is true!!")
		fallthrough
	case true:
		fmt.Println("Works here too because fallthrough!!")
	case true:
		fmt.Println("This not work!!!")
	default:
		fmt.Println("This does not work!")
	}

}