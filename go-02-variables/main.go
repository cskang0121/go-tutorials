// Topic : Variables

package main

import "fmt"

func main() {

	// Go is statically-typed language

	// 1. Explicitly-typed initialization
	var str1 string = "Kang Chin Shen"

	// 2. Implicitly-typed initialization
	var str2 = "Kang Chin Shen"

	// 3. Declaration (default to empty string)
	var str3 string

	// 4. Initialization without "var" (ONLY on creation)
	// - CANNOT use outside a function
	str4 := "Kang Chin Shen"

	// 5. Integer: int / int8 / int16 / int32 / int64 
	// 6. Unsigned integer: uint / unit8 / uint16 / uint32 / uint64 
	// 7. Float: float32 / float64

	// Test
	fmt.Println(str1, str2, str3, str4)
}