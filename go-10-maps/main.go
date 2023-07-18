// Topic : Maps

package main

import "fmt"

func main() {
	// 1. Map initialization: myMap := map[<key_type>]<value_type>{...}
	// 2. All keys must be of same type
	// 3. All values must be of same type
	myMap := map[int]string {
		1: "one",
		2: "two",
		3: "three",
	}

	// Loop through map
	for k, v := range myMap {
		fmt.Printf("%v : %v\n", k, v)
	}
}
