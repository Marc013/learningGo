package main

import (
	"fmt"

	"calc"
)

func init() {
	fmt.Println("")
	fmt.Println(">> Hello from main.go init() <<")
	fmt.Println("")
}

func main() {

	itemPrice := 50

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
