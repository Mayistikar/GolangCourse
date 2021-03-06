package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "martinis", "women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
		fmt.Printf("%T\n", v)
	}
}
