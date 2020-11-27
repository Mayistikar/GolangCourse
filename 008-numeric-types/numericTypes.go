package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

// Unsigned number, only positive numbers is available
var z uint8

func main() {

	// Asign a signed integer
	x = -22
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// Assing a float
	y = 34.2344
	// Truncated number is valid too
	// y = 34
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// Assing unsigned nunbers
	z = 255
	// z = 256 can not be, becaus overflow type size

	// Operative System Variable
	fmt.Println(runtime.GOOS)

	// OS Architecture
	fmt.Println(runtime.GOARCH)
}
