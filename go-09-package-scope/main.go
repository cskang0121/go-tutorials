// Scopes
// - Package-scoped variables 	: declared outside functions
// - Function-scoped variables : declared within functions

package main

import "fmt"

var mainVar int = 1

func printFileVar() {
	fmt.Println(fileVar)
}

// 1. Only 1 main function within 1 package
func main() {

	// 2. main.go accessing package-scoped variable in file.go -> print
	printFileVar()

	// 3. file.go accessing package-scoped variable in main.go
	printMainVar()
}

// 4. MUST specify all related files to run main(), e.g., go run main.go file.go