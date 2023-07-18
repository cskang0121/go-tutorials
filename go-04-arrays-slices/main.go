// Topic : Arrays & Slices

package main

import "fmt"

func main() { 
	// 1. Array – Fixed length
	// – Initialization
	var arr1 [3]int = [3]int{1, 2, 3} 	// Explicitly-typed
	var arr2 = [3]int{1, 2, 3}			// Implicitly-typed
	arr3 := [3]int{1, 2, 3}				// Without "var"

	// Test
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)
	
	// 1.2 Check length of an array
	fmt.Println("Length of arr1: ", len(arr1))

	// 1.3 Access values of an array
	fmt.Println("arr1[0]: ", arr1[0])

	// 2. Slices – Variable length
	var slc1 []int = []int{1, 2, 3} 	// Explicitly-typed
	var slc2 = []int{1, 2, 3}			// Implicitly-typed
	slc3 := []int{1, 2, 3}				// Without "var"

	// Test
	fmt.Println("slc1: ", slc1)
	fmt.Println("slc2: ", slc2)
	fmt.Println("slc3: ", slc3)

	// 2.2 Check length of a slice
	fmt.Println("Length of slc1: ", len(slc1))

	// 2.3 Access values of a slice
	fmt.Println("slc1[0]: ", slc1[0])

	// 2.4 Append to a slice
	slc1 = append(slc1, 4)

	fmt.Println("slc1: ", slc1)
}