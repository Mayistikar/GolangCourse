package main

import "fmt"

func main() {

	decimalNumber := 2
	fmt.Printf("Decimal: %d\tBinary: %b\n", decimalNumber, decimalNumber)

	// Bit shifting:
	otherNumber := decimalNumber << 1
	fmt.Printf("Decimal: %d\tBinary: %b\n", otherNumber, otherNumber)
}
