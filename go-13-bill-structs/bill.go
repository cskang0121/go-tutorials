package main

// 1. Define a struct: type <struct_name> struct {...}
type bill struct {
	name 	string
	items	map[string]float64
	tip		float64
}

// 2. Construct a struct object
// - Return type is the struct name
func newBill(name string) bill {
	return bill {
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
}