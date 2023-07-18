// Topic : Parsing Float

package main

import (
	"bufio"
	"os"		
	"strings"	
	"fmt"
	"strconv"
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

// 1. By default, struct is passed by value
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		// 2. strconv.ParseFloat() converts a string to a float of specified bits
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, priceFloat)

		fmt.Println("Item added - ", name, price)

		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount($): ", reader)
		fmt.Println(tip)

		tipFloat, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.updateTip(tipFloat)

		fmt.Println("Tip added - ", b.tip)

		promptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill", b)
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