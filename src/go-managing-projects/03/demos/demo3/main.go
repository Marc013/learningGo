package main

import (
	"calc"
	"fmt"
)

func main() {

	itemPrice := 50
	itemDiscount := 10
	var totalDiscount float32
	totalDiscount = calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
