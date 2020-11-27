package main

import "fmt"

type ownType int

var x ownType
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 123
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
