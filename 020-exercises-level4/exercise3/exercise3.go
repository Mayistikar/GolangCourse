package main

import (
	"fmt"
)

func main() {
	personalSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice1 := personalSlice[:5]
	fmt.Println(slice1)
	slice2 := personalSlice[5:]
	fmt.Println(slice2)
	slice3 := personalSlice[2:7]
	fmt.Println(slice3)
	fmt.Println(personalSlice[1:6])
}
