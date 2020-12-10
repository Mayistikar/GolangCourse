package main

import (
	"fmt"
)

func main() {
	x := make([]string, 5, 5)
	x = []string{"Apple", "Banana", "Coconut", "Pinneaple"}

	fmt.Println(x)
	fmt.Println("len: ", len(x), "cap:", cap(x))

	for i, v := range x {
		fmt.Println("idx:", i, "val:", v)
	}
}
