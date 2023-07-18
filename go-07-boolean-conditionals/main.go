// Topics : Boolean & Conditionals

package main

import "fmt"

func main() {
	// 1. Relational operators : >, >=, <, <=, ==. !=
	
	// 2. Conditionals
	// if <expression> {
	// } else if <expression> {
	// } else {
	// }

	// 3. continue & break
	for i := 0; i < 5; i++ {
		if i == 0 {
			fmt.Printf("pos %v : continue\n", i)
			continue
		} else if i > 3 {
			fmt.Printf("pos %v : break\n", i)
			break
		}
		fmt.Printf("pos %v\n", i)
	}
}