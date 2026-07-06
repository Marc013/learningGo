package main

import "fmt"

func init() {
	itemDiscount = 40
	fmt.Println("")
	fmt.Println(">> z.go init() itemDiscount =", itemDiscount)
	fmt.Println("")
}

func init() {
	itemDiscount = 50
	fmt.Println("")
	fmt.Println(">>AGAIN z.go init() itemDiscount =", itemDiscount)
	fmt.Println("")
}
