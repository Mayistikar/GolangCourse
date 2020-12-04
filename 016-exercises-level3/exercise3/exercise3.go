package main

import "fmt"

func main() {
	bornYear := 1990
	for bornYear <= 2020 {
		fmt.Println(bornYear)
		bornYear++
	}
}