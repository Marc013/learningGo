package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	itemSold := 10
	log.Println("itemSold is ", itemSold)
}
