package main

import "fmt"

func main () {

	var number int
	number = 42
	fmt.Println(number)
	fmt.Printf("%d\t%b\t%#x\n", number, number, number)
	shiftedNumber := number << 1
	fmt.Printf("%d\t%b\t%#x\n", shiftedNumber, shiftedNumber, shiftedNumber)
}