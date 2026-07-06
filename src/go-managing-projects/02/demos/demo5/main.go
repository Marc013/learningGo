package main

import (
	"fmt"

	"calc"
)

func main() {

	itemPrice := 50
	itemDiscount = disc100()

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
