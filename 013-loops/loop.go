package main

import "fmt"

func main () {

	/* FOR CLAUSE */
	fmt.Println("FOR CLAUSE")
	// for init statement; condition; post statement {}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// Nested loops
	fmt.Println("Nested Loops")
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			fmt.Print(y)
		}
		fmt.Println()
	}

	fmt.Println()

	/* CONDITION LOOP BLOCK*/
	fmt.Println("CONDITION LOOP BLOCK")
	count := 0
	for count < 10 {
		count ++
		fmt.Println("count: ",count)
	}

	fmt.Println()

	/* BREAK LOOP */
	fmt.Println("BREAK LOOP")
	numberRange := 10
	for {
		numberRange++
		if (numberRange == 20) {
			fmt.Println("Stop loop")
			break
		}
		fmt.Println(numberRange)
	}

	fmt.Println()

	/* CONTINUE JUMP*/
	fmt.Println("CONTINUE JUMP")
	counter := 0
	for {
		counter ++
		if (counter == 20) { break }
		if (counter%2 != 0) { continue }
		fmt.Println(counter)
	}

	fmt.Println()

	/* ASCII Printing */
	fmt.Println("ASCII Printing")
	ascii := 33;
	for {
		fmt.Printf("%#U\n", ascii)
		ascii++
		if (ascii == 123) { break }
	}

}