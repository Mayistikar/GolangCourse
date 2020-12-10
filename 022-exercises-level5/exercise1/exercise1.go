package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	person1 := person{
		first:    "Anderson",
		last:     "Rodriguez",
		iceCream: []string{"orange", "banana", "coconut"},
	}
	fmt.Println(person1)
	fmt.Println(person1.iceCream[0])
	for _, v := range person1.iceCream {
		fmt.Println("flavor:", v)
	}

	person2 := person{
		first:    "Wayra",
		last:     "Galindo",
		iceCream: []string{"chocolat", "orange"},
	}

	personMap := map[string]person{
		person1.last: person1,
		person2.last: person2,
	}

	fmt.Println(personMap["Rodriguez"])
	fmt.Println(personMap["Galindo"])

	for k, v := range personMap {
		fmt.Println("key:", k, "val:", v)
		for _, v := range v.iceCream {
			fmt.Println("ice:", v)
		}
	}
}
