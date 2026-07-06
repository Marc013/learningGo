package main

import (
	"calc"
	"fmt"

	"github.com/pioz/faker"
)

var itemDiscount = 0

func main() {

	itemPrice := 50
	itemDiscount := 50

	Curr := faker.CurrencySymbol()

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", Curr, itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", Curr, totalDiscount)

}
