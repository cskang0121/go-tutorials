// Topic : Hello World

// 1. Every Go file is part of a package
// - Naming the package "main" tells the compiler to convert the code into an executable
package main 

// 2. A package (Go Standard Library) for formatting string and printing messages
import "fmt"

func main() {

	// 3. Exported methods from packages start with uppercase
	fmt.Println("Hello World!")
}