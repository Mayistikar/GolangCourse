package main

import (
	"fmt"
)

// Struct definition
type person struct {
	first string
	last  string
	age   int
	mayor bool
}

// Embedding structs
type student struct {
	person
	active, live bool
}

func main() {

	// Value of struct
	person1 := person{
		first: "Anderson",
		last:  "Rodriguez",
		age:   30,
		mayor: true,
	}
	fmt.Println("Value of person1")
	fmt.Println(person1)
	fmt.Println(person1.first, person1.last)

	fmt.Println()

	person2 := person{
		first: "Wayra",
		last:  "Rodriguez",
		age:   24,
		mayor: true,
	}
	fmt.Println("Value of person2")
	fmt.Println(person2)
	fmt.Println(person2.first, person2.last)

	// Setting values for embedding structs
	student1 := student{
		person: person1,
		active: true,
		live:   true,
	}

	fmt.Println("Value of student1")
	fmt.Println(student1)

	fmt.Println(student1.person, student1.person.last)

	// Fields for inside struct are promoted to outside structs
	fmt.Println(student1.first, student1.last)

	fmt.Println()

	newStudent := student{
		person: person{
			first: "Camilo",
			last:  "Villalobos",
			age:   30,
			mayor: true,
		},
		active: false,
	}
	fmt.Println("Value of newStudent")
	fmt.Println(newStudent)

	// Composive literal
	differentStudent := student{
		person: person1,
	}
	fmt.Printf("%T\n", differentStudent)

	fmt.Println()

	// Anonymus Struct
	fmt.Println("Printing anonymus struct")
	anonymusStudent := struct {
		age   int
		mayor bool
	}{
		age:   34,
		mayor: true,
	}

	fmt.Println(anonymusStudent)

}
