//go:build solution

package task4

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	width, length := 3, 5
	expected := 15

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Errorf("wrong value for 'area', expected: %v, got %v", expected, area)
	}
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func TestCalculateAreaNegativeWidth(t *testing.T) {
	width, length := -4, 6
	expected := 0

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Errorf("wrong value for 'area', expected: %v, got %v", expected, area)
	}
	if err == nil {
		t.Error("expected an error")
	}
}

func TestCalculateAreaNegativeLength(t *testing.T) {
	width, length := 5, -7
	expected := 0

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Errorf("wrong value for 'area', expected: %v, got %v", expected, area)
	}
	if err == nil {
		t.Error("expected an error")
	}
}
