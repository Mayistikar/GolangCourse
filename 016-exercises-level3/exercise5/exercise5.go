package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println("Reminder modules: ", i%4)
	}
}