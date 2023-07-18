// Topic : "fmt" package

package main

import "fmt"

func main() {
	name := "Kang Chin Shen"
	age := 23

	// 1. fmt.Println()
	fmt.Println("My name is", name)

	// 2. fmt.Print()
	fmt.Print("My name is ", name, "\n")

	// 3. fmt.Printf() 
	// – Formatted string
	// – Format specifiers: 
		// – %v (default)
		// – %q (add double quotes - ONLY for string)
		// – %T (type of the variable)
		// – %f (float)
			// – %0.<x>f (<x> represents numbers of digits behind the floating point)
	fmt.Printf("My name is %v, and my age is %v \n", name, age)

	// 4. fmt.Sprintf() 
	// – Saved formatted string
	savedStr := fmt.Sprintf("My name is %v, and my age is %v", name, age)
	fmt.Println(savedStr)
}