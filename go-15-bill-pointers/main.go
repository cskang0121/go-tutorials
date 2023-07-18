// Topic : Struct Pointers 

package main

import "fmt"

func main() {
	myBill := newBill("Kang Chin Shen")

	myBill.addItem("coffee", 3.25)
	myBill.updateTip(10.00)

	fmt.Println(myBill.reportBill())
}