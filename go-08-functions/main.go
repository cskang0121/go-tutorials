// Topic : Loops

package main

import (
	"fmt"
	"strings"
)

// 1. Method signature : func <method_name> (<arg_1> <arg_type>) <return_type> {}
func printName(name string) {
	fmt.Println(name)
}

// 2. Passing a function as an argument
func cycleNames(names []string, f func(string)) {
	for _, val := range names {
		f(val)
	}
}

// 3. Return multiple values – enclose the return type with parentheses
func getInitials(name string) (string, string) {
	nameUpper := strings.ToUpper(name)

	nameStr := strings.Split(nameUpper, " ")

	var nameInitials []string

	for _, val := range nameStr {
		// 4. Use slicing for substring
		nameInitials = append(nameInitials, val[:1])
	}

	if len(nameInitials) < 2 {
		return nameInitials[0], "_"
	}

	return nameInitials[0], nameInitials[1]
}

func main() {
	names := []string{
		"Royston Lek Chun Keat", 
		"Kang Chin Shen", 
		// 5. '}' MUST BE next to the last element OR end with comma
		"Shaun Ting"} 

	cycleNames(names, printName)

	fn1, ln1 := getInitials("tifa lockhart")
	fn2, ln2 := getInitials("barret")

	fmt.Printf("%v %v\n", fn1, ln1)
	fmt.Printf("%v %v\n", fn2, ln2)
}