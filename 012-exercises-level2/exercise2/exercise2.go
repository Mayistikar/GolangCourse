package main

import "fmt"

func main() {
	g := 10 == 2
	h := 10 <= 2
	i := 10 >= 2
	j := 10 != 2
	k := 10 < 2
	l := 10 > 2

	fmt.Printf("g => value: %v\t type: %T\n", g, g)
	fmt.Printf("h => value: %v\t type: %T\n", h, h)
	fmt.Printf("i => value: %v\t type: %T\n", i, i)
	fmt.Printf("j => value: %v\t type: %T\n", j, j)
	fmt.Printf("k => value: %v\t type: %T\n", k, k)
	fmt.Printf("l => value: %v\t type: %T\n", l, l)
}
