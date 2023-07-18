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

func (b *bill) reportBill()string {
	fs := "Bill Breakdown:\n"
	
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-10v...$%v\n", k + ":", v)
		total += v
	}

	fs += fmt.Sprintf("%-10v...$%v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-10v...$%v", "total:", total + b.tip)

	return fs
}

// 1. Without *, receiver function creates a copy of the struct and update that copy
// 2. Adding * tells Go to update the value of the original copy
func (b *bill) updateTip(tip float64) {
	// 3. If the method signature indicates pass by reference, 
	// by default the struct object is pass by reference withint the function, 
	// NO de-referencing, e.g., (b*).tip = tip is required
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}