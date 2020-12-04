package main

import "fmt"

func main() {
	identifier := "favSoport"

	switch identifier {
	case "favSport":
		fmt.Println("Works!")
	case "favSoport":
		fmt.Println("Not works!")
	default:
		fmt.Println("Don't care!")
	}

}
