// Topic : "strings" & "sort" package

package main

// Import multiple packages
import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	greeting := "Hello World"

	// 1. "strings" package
	// 1.1 strings.Contains(<str>, <val>) returns true if <str> contains <val>, else returns false
	fmt.Println(strings.Contains(greeting, "Hello"))

	// 1.2 strings.ReplaceAll(<str>, <val>, <rplc>) replaces all <val> in <str> with <rplc>
	// – NOT inplace
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	// 1.3 strings.ToUpper(<str>) returns the string with all uppercase
	fmt.Println(strings.ToUpper(greeting))

	// 1.4 strings.ToLower(<str>) returns the string with all lowercase
	fmt.Println(strings.ToLower(greeting))

	// 1.5 strings.Index(<str>, <val>) returns the first found starting index of <val>, else returns -1
	fmt.Println(strings.Index(greeting, "ll"))

	// 1.6 strings.split(<str>, <delimiter>) returns an array of words by splitting the <str> using the <delimiter>
	fmt.Println(strings.Split(greeting, " "))


	// 2. "sort" package
	slc1 := []int{3, 2, 1}
	slc2 := []string{"coconut", "banana", "apple"}

	// 2.1 sort.Ints(<slc>) sorts the <slc> of integer-typed
	// – Inplace
	sort.Ints(slc1)
	fmt.Println(slc1)

	// 2.2 sort.SearchInts(<slc>, <val>) returns the first found index of <vaL>, else returns len(<slc>), i.e., max index of <slc> + 1
	// – Input <slc> must be sorted
	fmt.Println(sort.SearchInts(slc1, 3))

	// 2.3 sort.Strings(<arr>) sorts the <slc> of string-typed
	// – Inplace
	sort.Strings(slc2)
	fmt.Println(slc2)

	// 2.4 sort.SearchStrings(<arr>, <val>) returns the first found index of <vaL>, else returns len(<slc>), i.e., max index of <slc> + 1
	// – Input <slc> must be sorted
	fmt.Println(sort.SearchStrings(slc2, "coconut"))
}