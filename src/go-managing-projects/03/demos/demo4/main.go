package main

import (
	"fmt"

	"ps.m3.demo4/calc"
	//	"calc"
)

func main() {

	itemPrice := 50
	itemDiscount := 10

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
