// Topic : User Input

package main

import (
	"bufio"		// 1. Package to get user input
	"os"		// 2. Package to use os functionalities
	"strings"	
	"fmt"
)

// 3. r *bufio.Reader – r is a pointer 
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	// 4. bufio.NewReader - ReadString()
	// - Argument : a character that indicates end of reading
	// - Return Type : 
	// 		a. Input string
	// 		b. Error 
	inp, err := r.ReadString('\n')

	// 5. strings.Trimspace() takes an input string and return its trimmed version
	return strings.TrimSpace(inp), err
}

func createBill() bill { 
	// 6. bufio.NewReader() takes os.Stdin and returns a reader object
	reader := bufio.NewReader(os.Stdin)

	// 7. Put '_' if not interested in a particular input type
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill –", name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	fmt.Println(opt)
}	

func main() {
	myBill := createBill()

	promptOptions(myBill)

	fmt.Println(myBill)
}