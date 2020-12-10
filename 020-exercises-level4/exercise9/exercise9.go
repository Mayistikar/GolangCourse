package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "martinis", "women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(m)

	// Adding a key, value
	m["Anderson"] = []string{"Soccer", "Beer", "Code"}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Println()

	// Deleting a key, value
	delete(m, "Anderson")

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}
}
