// Topic : Switch Statement

package main

import (
	"bufio"
	"os"		
	"strings"	
	"fmt"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	inp, err := r.ReadString('\n')

	return strings.TrimSpace(inp), err
}

func createBill() bill { 
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill –", name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	// 1. Format of switch statement 
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s ...")
	default:
		fmt.Println("Invalid option ...")
		promptOptions(b)
	}
}	

func main() {
	myBill := createBill()

	promptOptions(myBill)

	fmt.Println(myBill)
}