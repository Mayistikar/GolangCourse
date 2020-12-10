package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, v := range slice {
		fmt.Println("idx:", i, "val:", v)
	}
	fmt.Printf("%T\n", slice)
}
