//go:build solution

package task2

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	width, length := 3, 5
	expected := 15

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestCalculateAreaNegativeWidth(t *testing.T) {
	width, length := -4, 6
	expected := 0

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
}

func TestCalculateAreaNegativeLength(t *testing.T) {
	width, length := 5, -7
	expected := 0

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
}
