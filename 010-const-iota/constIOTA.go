package main

import "fmt"

// Constants
const (
	a int     = 23
	b float64 = 23.4
	c string  = "Hello world"
)

// Iota => use to assign succesive untyped integer constants
const (
	x = iota
	y = iota
	z = iota
)

func main() {

	// Constants uses
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	// Obviusly a constant can not reassign
	// a = 22 is not allowed
	fmt.Println(a)

	// Iota uses
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
