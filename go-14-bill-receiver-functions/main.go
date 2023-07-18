// Topic : Receiver Functions

package main

import "fmt"

func main() {
	myBill := newBill("Kang Chin Shen")

	fmt.Println(myBill.reportBill())
}