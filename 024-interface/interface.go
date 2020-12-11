package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak(message string) {
	fmt.Println(p.first, "say", message)
}

// All struct that has a speak() function has a human type, polymorphism
type human interface {
	speak()
}

func run(p person) {
	fmt.Println(p.first, "is running")
}

func main() {
	fmt.Println("Interface implementation!")
	person1 := person{
		first: "Anderson",
		last:  "Rodriguez",
		age:   30,
	}

	person1.speak("Hola mundo desde persona1")

	run(person1)
}
