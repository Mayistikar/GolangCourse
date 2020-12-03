package main

import "fmt"

var number int = 42

func main() {
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Hexadecimal: %#x\n", number)
}
