package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooooo, James."}}

	fmt.Println(x)

	for _, v := range x {
		for _, w := range v {
			fmt.Printf("%v\t", w)
		}
		fmt.Println()
	}
}
