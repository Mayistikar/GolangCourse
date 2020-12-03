package main

import "fmt"

func main() {

	/* SIMPLE CONDITIONAL STATEMENT */
	fmt.Println("SIMPLE CONDITIONAL STATEMENT")
	if true {
		fmt.Println("Printing")
	}
	if false {
		fmt.Println("Not printed")
	}
	if !true {
		fmt.Println("Not printed")
	}
	if !false {
		fmt.Println("Printed")
	}

	fmt.Println()

	/* INITIALIZATION STATEMENT */
	fmt.Println("INITIALIZATION STATEMENT")
	if ifVariable := 23; ifVariable == 23 {
		fmt.Println("works, ifVariable => ", ifVariable)
		otherVariable := 21 // Have a if block scope
		var varVariable int = 12 // Have a if block scope
		const constVariable = 13 // Have a if block scope
		fmt.Println(otherVariable)
		fmt.Println(varVariable)
		fmt.Println(constVariable)
	}
	// "ifVariable" have a "if" scope, so ifVariable is not available out of "if" block
	// fmt.Println(ifVariable) is an sintax error
	// "otherVartiable" is not available in this scope
	// "varVariable" is not available in this scope
	// "constVariable" is not available in this scope
	
	fmt.Println()

	/* ELSE IF STATEMENT */
	fmt.Println("ELSE IF STATEMENT")
	var number int = 100
	if number == 40 {
		fmt.Println("Our Value is 40")
	} else if number == 41 {
		fmt.Println("Our value is 41")
	} else if number == 42 {
		fmt.Println("Our value is 42")
	} else {
		fmt.Println("Our value is not 40, 41 or 42")
	}

	fmt.Println()

	/* OPERATION CONDITIONAL STATEMENT */
	fmt.Println("OPERATION CONDITIONAL STATEMENT")
	otherNumber := 21
	if otherNumber/23 == 1 {
		fmt.Println("Other number is 23")
	} else if otherNumber%2 == 0{
		fmt.Println("Other number is even")
	} else if otherNumber*2 == 42 {
		fmt.Println("Other number is 21")
	}

	fmt.Println()

	/* CONDITIONAL LOGIC OPERATORS */
	fmt.Println("CONDITIONAL LOGIC OPERATORS")
	fmt.Println("true && true => ", true && true)
	fmt.Println("true && false => ", true && false)
	fmt.Println("false && false => ", false && false)
	fmt.Println("true || true => ", true || true)
	fmt.Println("true || false => ", true || false)
	fmt.Println("false || false => ", false || false)
	fmt.Println("!true => ", !true)
	fmt.Println("!false => ", !false)

}