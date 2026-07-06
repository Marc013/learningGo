package main

import (
	"fmt"

	"calc"
)

var itemDiscount = 0

func init() {
	itemDiscount = 30
	fmt.Println("")
	fmt.Println(">> main.go init() itemDiscount =", itemDiscount)
	fmt.Println("")
}

func main() {

	itemPrice := 50
	itemDiscount = 75
	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
