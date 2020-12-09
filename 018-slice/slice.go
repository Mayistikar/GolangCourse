package main

import "fmt"

func main() {

	/* COMPOSITE LITERAL */
	fmt.Println("COMPOSITE LITERAL")
	// Slice allows group together values of the same type
	numberSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numberSlice)
	// Getting length of slice
	fmt.Println(len(numberSlice))
	// Getting index position value
	fmt.Println(numberSlice[3])
	strignSlice := []string{"Hello", "Slice", "World"}
	fmt.Println(strignSlice)

	fmt.Println()

	/* WALKING OVER SLICE */
	fmt.Println("WALKING OVER SLICE")
	fmt.Println(numberSlice)
	for i, v := range numberSlice {
		fmt.Println("idx:", i, " => val:", v)
	}

	fmt.Println()

	/* PRINTING RANGE OF SLICE */
	fmt.Println("PRINTING RANGE OF SLICE")
	fmt.Println(numberSlice)
	fmt.Println(numberSlice[1:2])
	fmt.Println(numberSlice[1:3])
	fmt.Println(numberSlice[:len(numberSlice)])

	fmt.Println()

	/* APPEND SLICES */
	fmt.Println("APPEND SLICES")
	normalSlice := []int{2, 3, 4, 5}
	fmt.Println(normalSlice)
	appendedSlice := append(normalSlice, 10, 11, 12)
	fmt.Println("Appended: ", appendedSlice)
	// Append two slice, using VARIADIC PARAMETER
	otherAppendedSlice := append(normalSlice, appendedSlice...)
	fmt.Println(otherAppendedSlice)

	fmt.Println()

	/* DELETE SLICE ELEMENT */
	fmt.Println("DELETE SLICE ELEMENT")
	sliceExample := []string{"Apple", "Banana", "Coconut", "Orange"}
	fmt.Println(sliceExample)
	// Deleting a "Coconut" fruit
	sliceExample = append(sliceExample[0:2], sliceExample[3:]...)
	fmt.Println(sliceExample)

	fmt.Println()

	/* MAKE SLICE */
	fmt.Println("MAKE SLICE")
	// make(type, length, capacity)
	makingSlice := make([]int, 10, 100)
	fmt.Println(makingSlice)
	fmt.Println("length", len(makingSlice))
	fmt.Println("capacity", cap(makingSlice))

	fmt.Println()

	/* MULTIDIMENSIONAL SLICES */
	fmt.Println("MULTIDIMENSIONAL SLICES")
	binarySlice := []int{1, 0, 0, 1, 1, 1, 0}
	decimalSlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySlice, decimalSlice)
	// Multidimensional
	multiSlice := [][]int{binarySlice, decimalSlice}
	fmt.Println(multiSlice)
}
