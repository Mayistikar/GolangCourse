package main

import "fmt"

func main() {
	s := `"Hello World!!!
	
		from mars!!!
	"`
	s = "hola mundo"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// A string represent a slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

}
