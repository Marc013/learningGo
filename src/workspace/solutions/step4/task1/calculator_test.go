//go:build solution

package task1

import (
	"os"
	"testing"
)

var c Calculator

func TestMain(m *testing.M) {
	c = Calculator{}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestAdd(t *testing.T) {
	expected := 5.0

	result := c.Add(2, 3)

	if result != expected {
		t.Errorf("Addition was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	expected := 2.0

	result := c.Subtract(5, 3)

	if result != expected {
		t.Errorf("Subtraction was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	expected := 6.0

	result := c.Multiply(2, 3)

	if result != expected {
		t.Errorf("Multiplication was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestDivideValid(t *testing.T) {
	expected := 2.0

	result, err := c.Divide(6, 3)

	if err != nil {
		t.Errorf("Error while dividing: %v", err)
	}

	if result != expected {
		t.Errorf("Division was incorrect, got: %f, want: %f.", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := c.Divide(6, 0)

	if err == nil {
		t.Error("Expected error for division by zero but got none")
	}
}
