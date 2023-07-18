// Topic : Pass By Reference

package main

// For datatypes default to pass by value:
// - String
// - Int
// - Float
// - Boolean
// - Array
// - Struct
// - There are 2 ways to update the value:
	// a. Return value to the caller
	// b. Pass by reference

import "fmt"

// a. Return value to the caller
func func1(x int) int{
	return x + 1
}

// b. Pass by reference
func func2(x *int) {
	*x += 1
}

func main() {
	name := "Kang Chin Shen"

	// 1. Retrieve the memory address of a variable
	ptr := &name
	fmt.Println("ptr :", ptr)
	// 2. De-reference value of a memory address
	fmt.Println("de-referencing ptr:", *ptr)

	
	x := 1
	// a. Return value to the caller
	x = func1(x)
	fmt.Println("x:", x)

	// b. Pass by reference
	func2(&x)
	fmt.Println("x:", x)
}