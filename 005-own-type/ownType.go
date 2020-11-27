package main

import "fmt"

// Define a new type call "myType" underline int
type myType int

var myTypeVariable myType
var intVariable int

func main() {
	myTypeVariable = 23
	fmt.Println(myTypeVariable)
	// Print the variable's type to show "main.myType"
	fmt.Printf("%T\n", myTypeVariable)

	// In Go there is not CASTING, we talking about CONVERSION
	intVariable = int(myTypeVariable)
	fmt.Println(intVariable)
	fmt.Printf("%T\n", intVariable)
}
