// Topic : Loops

package main

import "fmt"

func main() {
	// 1. While-loop
	x := 1
	for x < 5{
		fmt.Println("x :", x)
		x++
	}

	// 2. For-loop
	for x:=1; x < 5; x++ {
		fmt.Println("x :", x)
	}

	// 3. Loop through an array / slice with index only
	arr := [3]string{"a", "b", "c"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 4. Loop through an array / slice with value only
	for _, val := range arr {
		fmt.Printf("The value is %v.\n", val)
	}	

	// 5. Loop through an array / slice with both index and value
	for idx, val := range arr {
		fmt.Printf("The value at position %v is %v.\n", idx, val)
	}
}