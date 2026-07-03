package task2

import (
	"errors"
	"fmt"
)

type Calculator struct{}

func (c Calculator) Add(a, b float64) float64 {
	fmt.Printf("Add called with parameters %v and %v\n", a, b)
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	fmt.Printf("Subtract called with parameters %v and %v\n", a, b)
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	fmt.Printf("Multiply called with parameters %v and %v\n", a, b)
	return a * b
}

func (c Calculator) Divide(a, b float64) (float64, error) {
	fmt.Printf("Divide called with parameters %v and %v\n", a, b)
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
