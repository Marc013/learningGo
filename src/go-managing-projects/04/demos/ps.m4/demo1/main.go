package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed numbers.txt
	data []byte
)

func main() {
	fmt.Println(string(data))

	fmt.Println("----------------------")

	lines := strings.Split(string(data), "\r\n")
	var sum int
	for _, line := range lines {
		if line != "" {
			val, _ := strconv.Atoi(line)
			sum += val
		}
	}
	fmt.Println("Sum of all numbers: ", sum)
}
