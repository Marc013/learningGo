package main

import (
	"fmt"

	"ps.m3.demo5/calc"
)

func main() {

	itemPrice := 50
	itemDiscount = disc100()
	disc50()

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
