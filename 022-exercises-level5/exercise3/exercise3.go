package main

import (
	"fmt"
)

func main() {
	anonymusName := struct {
		first   string
		last    string
		age     int
		friends map[string]struct {
			first string
			last  string
			phone string
		}
	}{
		first: "Anderson",
		last:  "Rodriguez",
		age:   30,
		friends: map[string]struct {
			first string
			last  string
			phone string
		}{
			"Leo": {
				first: "Leonardo",
				last:  "Fernandez",
				phone: "31180309273",
			},
		},
	}
	fmt.Println(anonymusName)
	fmt.Println(anonymusName.friends["Leo"])
}
