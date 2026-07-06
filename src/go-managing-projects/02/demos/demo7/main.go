package main

import (
	"defaults"
	"fmt"

	"calc"
)

func main() {

	itemPrice := 50
	itemDiscount = defaults.ItemDiscountValue
	//itemDiscount = defaults.itemDiscountValue

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
