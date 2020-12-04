package main

import "fmt"

func main() {
	counter := 1
	for {
		fmt.Println(counter)
		if (counter == 10000) { break }
		counter++
	}
}