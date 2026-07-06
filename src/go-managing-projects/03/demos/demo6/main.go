package main

import (
	"fmt"

	"github.com/pioz/faker"
	"ps.m3.demo6/calc"
)

func main() {

	itemPrice := 50
	itemDiscount := 50

	Curr := faker.CurrencySymbol()

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", Curr, itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", Curr, totalDiscount)

}
