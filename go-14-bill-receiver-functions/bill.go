package main

import "fmt"

type bill struct {
	name 	string
	items	map[string]float64
	tip		float64
}

func newBill(name string) bill {
	return bill {
		name: name,
		items: map[string]float64{
			"pie": 5.99,
			"cake": 3.99,
		},
		tip: 0,
	}
}

// 1. Add (<name> <type>) in front of a function to make it a receiver function of the <type>
func (b bill) reportBill()string {
	fs := "Bill Breakdown:\n"
	
	var total float64 = 0

	for k, v := range b.items {
		// 2. "+" concatenates strings
		// 3. Add positive or negative integer between %v to specify the reserved space
		fs += fmt.Sprintf("%-10v...$%v\n", k + ":", v)
		total += v
	}

	fs += fmt.Sprintf("%-10v...$%v", "total:", total)

	return fs
}