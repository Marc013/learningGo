package task3

import (
	"errors"
	"fmt"
)

func CalculateArea(width, length int) (int, error) {
	fmt.Printf("CalculateArea called with parameters %v and %v\n", width, length)

	if width <= 0 {
		return 0, errors.New("width is too low")
	}

	return width * length, nil
}
