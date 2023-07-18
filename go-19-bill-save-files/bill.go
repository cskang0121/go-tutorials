package main

import (
	"fmt"
	"os"
)

type bill struct {
	name 	string
	items	map[string]float64
	tip		float64
}

func newBill(name string) bill {
	return bill {
		name: name,
		items: map[string]float64{},
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

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	// 1. MUST convert the files to slice of bytes before saving it
	// 2. byte() convert a string into bytes
	data := []byte(b.reportBill())

	// 3. os.WriteFiles writes the data in bytes into a file
	// Arguments:
	// 		a. file path
	// 		b. data of slice of bytes
	// 		c. permissions – 0644 
	// Returns error if any
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		// 4. panic() takes in an error, stop the program and display error
		panic(err)
	}
	fmt.Println("bill was saved to file")
}