package main

import "fmt"

/*
	Declaration using "var" keyword, allow use the variable in a
	PACKAGE scope
*/
var outVariable = "Program scope"

/*
	With var declaration, can define a variable with ZERO VALUE default assign
	ZERO VALUE => is a default assign that depends of variable type:
	false for booleans, 0 for integers, 0.0 for floats, "" for strings,
	and nil for pointers, functions, interfaces, slices, channels and maps.
*/
var booleanVariable bool
var integerVariable int
var floatVariable32 float32
var floatVariable64 float64
var stringVariable string

func main() {

	fmt.Println(outVariable)

	/*
		Declaration using := just allow a intern scope
	*/
	inVariable := "Function scope"

	fmt.Println(inVariable)

	// ZERO VALUE examples:
	fmt.Println("booleanVariable: ", booleanVariable)
	fmt.Println("integerVariable: ", integerVariable)
	fmt.Println("floatVariable: ", floatVariable32)
	fmt.Println("floatVariable64: ", floatVariable64)
	fmt.Println("stringVariable: ", stringVariable)

	/*
		When declared and assigned a variable, the variable
		takes the value type
	*/
	var stringTypeVariable = "String type value"
	fmt.Println(stringTypeVariable)
	// Printing variable's type
	fmt.Printf("%T\n", stringTypeVariable)

	/*
		GO is a STATIC PROGRAMMING LANGUAGE, so you can't redefine a variable's types
	*/
	var integerTypeVariable = 100
	// Can not change the integerTypeVariable type.
	// integerTypeVariable = "string value" is forbidden
	fmt.Printf("%T\n", integerTypeVariable)

}
