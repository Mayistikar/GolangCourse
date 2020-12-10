package main

import (
	"fmt"
)

func main() {

	holds := [5]int{1, 2, 3, 4, 5}

	for idx, v := range holds {
		fmt.Println("Index:", idx, "Value:", v)
	}
	fmt.Printf("%T\n", holds)

}
