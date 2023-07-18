// Topic : Pass By Values & Pass By Reference
// - Pass By Values
	// - String
	// - Int
	// - Float
	// - Boolean
	// - Array
	// - Struct
// - Pass By Reference
	// - Slice
	// - Map
	// - Functions

package main

import "fmt"

// 1. Go creates a copy of x (in main) and assigns it to x (in the local scope of the function)
// 2. Go updates the value of the x (in the local scope of the function)
func passByValue(x int) {
	x = 2
}

// 3. Go creates a copy of pointer (points to y in main) and assigns it to y (in the local scope of the function)
// 4. Go updates the value of the underlying structure pointed by y (in the local scope of the function)
func passByReference(y map[int]string) {
	y[4] = "four"
}

func main() {
	x := 1
	passByValue(x)
	fmt.Println("x:", x)

	y := map[int]string {
		1 : "one",
		2 : "two",
		3 : "three",
	}
	passByReference(y)
	fmt.Println(y)
}