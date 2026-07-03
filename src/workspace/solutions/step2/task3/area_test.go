//go:build solution

package task3

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
