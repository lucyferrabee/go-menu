package main

import "fmt"

func main() {
	mybill := newBill("Mario's bill")

	mybill.updateTip(10)
	mybill.addItem("pizza", 8.75)
	mybill.addItem("onion soup", 4.99)
	mybill.addItem("macaroni cheese", 10.99)
	mybill.addItem("jacket potato", 6.50)

	fmt.Println(mybill.format())
}
