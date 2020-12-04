package main

import "fmt"

func main() {

	bornYear := 1990
	for {
		if (bornYear > 2020) { break }
		fmt.Println(bornYear)
		bornYear++
	}
}